#### <sub><sup><a name="4583" href="#4583">:link:</a></sup></sub> feature

* We've introduced a `components` table in order to better synchronize all the internal processes that run on the web nodes.

This should help reduce the amount of duplicated work (when running more than 1 ATC), and decrease the load on your database.

There is no configuration required to take advantage of these new improvements.

#### <sub><sup><a name="4535" href="#4535">:link:</a></sup></sub> feature, breaking

* The CloudFoundry auth connector, when configured to authorize users based on CF space membership, will now authorize space auditors and space managers in addition to space developers. This is a breaking change as any teams with CF space-based configuration may grant access to users that they wouldn't have before. #4661

#### <sub><sup><a name="4480" href="#4480">:link:</a></sup></sub> feature

* @ProvoK added support for a `?title=` query parameter on the pipeline/job badge endpoints! Now you can make it say something other than "build". #4480
  ![badge](https://ci.concourse-ci.org/api/v1/teams/main/pipelines/concourse/badge?title=check%20it%20out)

#### <sub><sup><a name="4518" href="#4518">:link:</a></sup></sub> feature

* @evanchaoli added a feature to stop ATC from attempting to renew Vault leases that are not renewable #4518.

#### <sub><sup><a name="4516" href="#4516">:link:</a></sup></sub> feature

* Add 5 minute timeout for baggageclaim destroy calls #4516.

#### <sub><sup><a name="4334" href="#4334">:link:</a></sup></sub> feature

* @aledeganopix4d added a feature sort pipelines alphabetically #4334.

#### <sub><sup><a name="4470" href="#4470">:link:</a></sup></sub> feature, breaking

* All API payloads are now gzipped. This should help save bandwidth and make the web UI load faster. #4470

#### <sub><sup><a name="4494" href="#4494">:link:</a></sup></sub> feature

* API endpoints have been changed to use a single transaction per request, so that they become "all or nothing" instead of holding data in memory while waiting for another connection from the pool. This could lead to snowballing and increased memory usage as requests from the web UI (polling every 5 seconds) piled up. #4494

#### <sub><sup><a name="4448-4588" href="#4448-4588">:link:</a></sup></sub> feature

* You can now pin a resource to different version without unpinning it first #4448, #4588.

#### <sub><sup><a name="4507" href="#4507">:link:</a></sup></sub> fix

* @iamjarvo fixed a [bug](444://github.com/concourse/concourse/issues/4472) where `fly builds` would show the wrong duration for cancelled builds #4507.

#### <sub><sup><a name="4590" href="#4590">:link:</a></sup></sub> feature

* @pnsantos updated the Material Design icon library so now the `concourse-ci` icon is available for resources :tada: #4590

#### <sub><sup><a name="4492" href="#4492">:link:</a></sup></sub> fix

* The `fly format-pipeline` now always produces a formatted pipeline, instead of declining to do so when it was already in the expected format. #4492

#### <sub><sup><a name="3600" href="#3600">:link:</a></sup></sub> feature

* Concourse now garbage-collects worker containers and volumes that are not tracked in the database. In some niche cases, it is possible for containers and/or volumes to be created on the worker, but the database (via the web) assumes their creation had failed. If this occurs, these untracked containers can pile up on the worker and use resources. #3600 ensures that they get cleaned appropriately.

#### <sub><sup><a name="4606" href="#4606">:link:</a></sup></sub> feature

* @wagdav updated worker heartbeat log level from `debug` to `info` to reduce extraneous log output for operators #4606

#### <sub><sup><a name="4625" href="#4625">:link:</a></sup></sub> fix

* Fixed a [bug](https://github.com/concourse/concourse/issues/4313) where your dashboard search string would end up with `+`s instead of spaces when logging in. #4265

#### <sub><sup><a name="4636" href="#4636">:link:</a></sup></sub> fix

* Fixed a [bug](https://github.com/concourse/concourse/issues/4493) where the job page would show a loading spinner forever when there were no builds (like before the job had ever been run) #4636.

#### <sub><sup><a name="4630" href="#4630">:link:</a></sup></sub> fix

* Fixed a [bug](https://github.com/concourse/concourse/issues/3921) where the tooltip that says 'new version' on a get step on the build page could be hidden underneath the build header #4630.

#### <sub><sup><a name="4637" href="#4637">:link:</a></sup></sub> fix

* Fixed a [bug](https://github.com/concourse/concourse/issues/3942) where log lines on the build page would have all their timestamps off by one. #4637

#### <sub><sup><a name="4668" href="#4668">:link:</a></sup></sub> fix

* @evanchaoli fixed a [bug](https://github.com/concourse/concourse/issues/4656) where secret redaction incorrectly "redacts" empty string resulting in mangled logs. #4668
