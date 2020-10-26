# SDK Examples

This repository has examples for common tasks across all SDKs.
It includes things like basic connecting, bulk loading, N1QL queries with placeholders (a.k.a. parameters), counters, Sub-Document operations and so on.
While the primary source of documentation is the [SDK Docs][2], it is often useful to have a clone-able, editable and runnable sample.
That is what this repository attempts to provide.
Most samples will be available in multiple lanugages allowing a developer to see how it's done in one language versus another.


## Development and Branching

At the moment, there is only the `master` branch.
These samples represent the "SDK 3.0" (in quotes because it refers to API level, not the specific SDK version) compatible code samples.

The samples contained here are developed in one of two ways.
The primary way is that a code sample is authored as part of the [SDK Documentation Repo][1].
It need not necessarily appear in the documentation, but it frequently will.
Those samples are copied, with minimal processing to remove the Antora tags, to this repository.
Typically, this is done as a script to a 'peer directory' and then separately committed.
To keep things simple, we do not try to leverage submodules or the like.

The second way is that a code sample maybe directly added to this repository.
This is less preferred as it makes the use of the sample in the SDK repository impossible and could lead to a duplicate sample.
It may be appropriate when the sample is very unique to something that is not and will not likely be in the documentation.


[1]: <https://github.com/couchbase/?q=docs-sdk&type=&language=> "All SDK Docs Repos"
[2]: <https://docs.couchbase.com/home/sdk.html> "SDK Documentation on docs.couchbase.com"