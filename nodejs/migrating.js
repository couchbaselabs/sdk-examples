const couchbase = require ("couchbase");
const RawBinaryTranscoder = require("./rawbinarytranscoder");


function start(){
 return new Promise( (resolve, reject) => { resolve(); });
}

var cluster;
var bucket;
var collection;

start()
  .then(connect)
  //.then(connstr)
  .then(simpleget)
  .then(upsertandget)
  .then(rawbinary)
  .then(customtimeout)
  .then(querysimple)
  // .then(analyticssimple) // don't have dataset
  // .then(analyticsparameterized) // don't have dataset
  .then(searchsimple) 
  .then(searchcheck)
  // .then(viewquery)
  .then(shutdown)
  .then(process.exit)



async function timeoutbuilder_java_only(){
      // SDK API 3 equivalent
      const env = ClusterEnvironment
        .builder()
        .timeoutConfig(TimeoutConfig.kvTimeout(Duration.ofSeconds(5)))
        .build();
}

async function connect(){
  const options = { username:"Administrator", password:"password"};
  cluster = await couchbase.connect( "http://127.0.0.1", options);
  bucket = cluster.bucket("travel-sample");
  collection = bucket.scope("inventory").collection("airport");
}

async function sysprops_java_only(){
/*** this only has an effect for java ***/
      // Will set the max http connections to 23
      process.env["com.couchbase.env.io.maxHttpConnections"]="23"
      console.log( process.env["com.couchbase.env.io.maxHttpConnections"]);
      // This is equivalent to
      //ClusterEnvironment env = ClusterEnvironment
      //  .builder()
      //  .ioConfig(IoConfig.maxHttpConnections(23))
      //  .build();
}

async function connstr(){
      if(cluster) cluster.close();
      const options = { username:"Administrator", password:"password"};
      cluster = await couchbase.connect( "http://127.0.0.1/?query_timeout=2000", options);
}

/*
    {
      Cluster.connect("127.0.0.1", "username", "password");
    }

    {
      Cluster.connect(
        "127.0.0.1",
        clusterOptions(PasswordAuthenticator.create("username", "password"))
      );
    }
*/

async function certauth(){ // use one in  connecting-cert-auth.js
}

async function simpleget() {
  console.log("simpleget:");
  
  // tag::simpleget[]
  
  const getResult = await collection.get("airport_1254");
  
  // end::simpleget[]
  console.log(getResult);
}


async function upsertandget(){
      console.log("upsertandget:");
      const  upsertResult = await collection.upsert("mydoc-id", { myvalue: "me"});
      const getResult = await collection.get("mydoc-id");
      console.log(getResult);
    }

async function rawbinary(){
      console.log("rawbinary:");
      const content = Buffer.from("some data to become binary");
      const  upsertResult = await collection.upsert(
        "mydoc-id",
        content,
        {transcoder:new RawBinaryTranscoder()}
      );
      const getResult = await collection.get("mydoc-id");
      console.log(getResult);
}

async function customtimeout(){
      console.log("customtimeout:");
      // SDK API 3 custom timeout
      const getResult = await collection.get( "airport_1254", {timeout : 2000});
      console.log(getResult);
}

async function querysimple(){
      console.log("querysimple :");
      // SDK API 3 simple query
      const queryResult = await cluster.query("SELECT * FROM `travel-sample` WHERE city=$1 LIMIT 10", { parameters: ['Paris']});
      queryResult.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function queryparameterized(){
      console.log("queryparameterized-named/positional :");
      // SDK API 3 named parameters
      const queryResultNamed = await cluster.query("SELECT * FROM `travel-sample` WHERE city=$CITY LIMIT 10", { parameters: {CITY:'Paris'}});
      queryResultNamed.rows.forEach((row)=>{
         console.log(row);
      });
      // SDK API 3 positional parameters
      const queryResultPositional = await cluster.query("SELECT * FROM `travel-sample` WHERE city=$1 LIMIT 10", { parameters: ['Paris']});
      queryResultPositional.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function analyticssimple(){
      console.log("analyticssimple :");
      // SDK API 3 simple analytics query
      const analyticsResult = await cluster.analyticsQuery("select * from `travel-dataset` LIMIT 10");
      analyticsResult.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function analyticsparameterized(){
      console.log("analyticsparameterized-named/positional :");
      // SDK API 3 named parameters for analytics
      const analyticsResult1 = await cluster.analyticsQuery(
        "select * from `travel-sample` where city = $CITY LIMIT 10",
        { parameters : { CITY : "Paris"}}
      ).catch((e)=>{console.log(e); throw e;});
      analyticsResult1.rows.forEach((row)=>{
         console.log(row);
      });

      // SDK API 3 positional parameters for analytics
      const analyticsResult2 = await cluster.analyticsQuery(
        "select * from `travel-sample` where city = $1 LIMIT 10",
        { parameters : [ "airport" ] }
      );
      analyticsResult2.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function searchsimple(){
      console.log("searchsimple :");
      // SDK API 3 search query
      const ftsQuery =  couchbase.SearchQuery.match("airport");
      const searchResult = await cluster.searchQuery(
        "hotels",
        ftsQuery,
        { timeout:2000,
          limit:5,
          fields : ["a", "b", "c"] },
        (err, res) => {
		if(err) console.log(err);
		if(res) console.log(res);
        }
      );
      searchResult.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function searchcheck(){
      console.log("searchcheck :");
      const ftsQuery =  couchbase.SearchQuery.match("airport");
      const searchResult = await cluster.searchQuery(
        "hotels",
        ftsQuery,
        { timeout:2000,
          limit:5}
      ).catch((e)=>{console.log(e); throw e;});
      console.log(searchResult);
      if (searchResult.meta.status.failed == 0) {
        searchResult.rows.forEach((row)=>{
           console.log(row);
        });
      }
    }

async function viewquery(){
      console.log("viewquery :");
      // SDK API 3 view query
      const viewResult = await bucket.viewQuery(
        "dev_airports",
        "airport_view",
        { limit:5, skip:2, timeout:10000 }
      ).catch((e)=>{console.log(e); throw e;});
      viewResult.rows.forEach((row)=>{
         console.log(row);
      });
    }

async function shutdown(){
  console.log("shutdown");
  if (cluster && cluster.name && cluster.name == 'Error') {
    throw cluster;
  }
  if (cluster) {
  // tag::shutdown[]
  cluster.close();
  // end::shutdown[]
  } else {
    console.log(new Error("cluster undefined - never opened"));
  }
  cluster = new Error("cluster already closed here");
}
