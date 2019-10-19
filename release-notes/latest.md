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

#### <sub><sup><a name="4448-4588" href="#4448-4588">:link:</a></sup></sub> feature

* You can now pin a resource to different version without unpinning it first #4448, #4588.

#### <sub><sup><a name="4507" href="#4507">:link:</a></sup></sub> fix

* @iamjarvo fixed a [bug](444://github.com/concourse/concourse/issues/4472) where `fly builds` would show the wrong duration for cancelled builds #4507.

#### <sub><sup><a name="4590" href="#4590">:link:</a></sup></sub> feature

* @pnsantos updated the Material Design icon library so now the `concourse-ci` icon is available for resources :tada: #4590

#### <sub><sup><a name="4492" href="#4492">:link:</a></sup></sub> fix

* The `fly format-pipeline` now always produces a formatted pipeline, instead of declining to do so when it was already in the expected format. #4492

#### <sub><sup><a name="4617" href="#4617">:link:</a></sup></sub> feature

* New Relic emitter logs info message when API call to New Relic responds with a non 2xx status code.  #4617