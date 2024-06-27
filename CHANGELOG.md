# Changelog

## 1.0.0 (2024-06-27)


### Features

* add backup service ([#19](https://github.com/cloudnative-pg/cnpg-i/issues/19)) ([92ca0ea](https://github.com/cloudnative-pg/cnpg-i/commit/92ca0eaa10ad0dc88b7458b56331092453e9035f))
* add lifecycle hooks ([#20](https://github.com/cloudnative-pg/cnpg-i/issues/20)) ([5869c64](https://github.com/cloudnative-pg/cnpg-i/commit/5869c644788b8d4d3de0b5305bd0fc3c7f1be17b))
* add operator service type in the plugin metadata ([#6](https://github.com/cloudnative-pg/cnpg-i/issues/6)) ([7fcdddb](https://github.com/cloudnative-pg/cnpg-i/commit/7fcdddbaf7f46c89affcf3871ddd9fb51e691adc))
* add plugin metadata ([#7](https://github.com/cloudnative-pg/cnpg-i/issues/7)) ([e16cc54](https://github.com/cloudnative-pg/cnpg-i/commit/e16cc543b9bb875b02d10b4cfa015ebd3050ea22)), closes [#4](https://github.com/cloudnative-pg/cnpg-i/issues/4)
* add pre and post reconcile hooks ([#21](https://github.com/cloudnative-pg/cnpg-i/issues/21)) ([af27643](https://github.com/cloudnative-pg/cnpg-i/commit/af2764325710cd98881e211810d2a01f61b62d43))
* operator reconciler specification ([#3](https://github.com/cloudnative-pg/cnpg-i/issues/3)) ([0b01870](https://github.com/cloudnative-pg/cnpg-i/commit/0b018708a768a8f4d59e99716a782e44cceb7ff2))
* pluggable dynamic controller ([#25](https://github.com/cloudnative-pg/cnpg-i/issues/25)) ([7d938f6](https://github.com/cloudnative-pg/cnpg-i/commit/7d938f60f998af1d336a669a2497d1cf68b9d582))
* remove MutatePod method ([#26](https://github.com/cloudnative-pg/cnpg-i/issues/26)) ([9afc581](https://github.com/cloudnative-pg/cnpg-i/commit/9afc58113460f207ca4fb83b6ba875caef1abb64))
* service for management of WAL files ([#2](https://github.com/cloudnative-pg/cnpg-i/issues/2)) ([963073d](https://github.com/cloudnative-pg/cnpg-i/commit/963073d61df217fd6c2cd9f4ca77b4a57e623c04)), closes [#1](https://github.com/cloudnative-pg/cnpg-i/issues/1)


### Bug Fixes

* add cluster definition to MutatePod RPC call ([#10](https://github.com/cloudnative-pg/cnpg-i/issues/10)) ([2492dba](https://github.com/cloudnative-pg/cnpg-i/commit/2492dba8ebbd98f1da02a9e59d05e5a76beac8e6))
* **deps:** update all non-major go dependencies ([caf6b1b](https://github.com/cloudnative-pg/cnpg-i/commit/caf6b1b23daa7a733ad2a6c321ab097660bbdd94))
* package names ([#30](https://github.com/cloudnative-pg/cnpg-i/issues/30)) ([b0b3107](https://github.com/cloudnative-pg/cnpg-i/commit/b0b310788fa1c097e139e31d661f124be91b94b7))
* reconciler package path ([#27](https://github.com/cloudnative-pg/cnpg-i/issues/27)) ([3624915](https://github.com/cloudnative-pg/cnpg-i/commit/362491594dde7a5360b24a0fd14c7e02715cab09))
* regenerate stub & skeleton ([#11](https://github.com/cloudnative-pg/cnpg-i/issues/11)) ([bc5f2ca](https://github.com/cloudnative-pg/cnpg-i/commit/bc5f2cab5e5949cc4e6dd0d2200613aea7330f20))
* remove maintainers from plugin metadata ([#12](https://github.com/cloudnative-pg/cnpg-i/issues/12)) ([a18f8f5](https://github.com/cloudnative-pg/cnpg-i/commit/a18f8f50cf36d1928a5d41ebeba192228cf6d2ea))
* remove stale constants ([#28](https://github.com/cloudnative-pg/cnpg-i/issues/28)) ([4adc035](https://github.com/cloudnative-pg/cnpg-i/commit/4adc03536d626577c919718d7d02bef89fbe366a))
* service capability order ([#14](https://github.com/cloudnative-pg/cnpg-i/issues/14)) ([bc73ed3](https://github.com/cloudnative-pg/cnpg-i/commit/bc73ed38c383c5e2b90c95ccb67a724e8354d5e7))
* use bytes for JSON serialization instead of strings ([#13](https://github.com/cloudnative-pg/cnpg-i/issues/13)) ([e99be6d](https://github.com/cloudnative-pg/cnpg-i/commit/e99be6d0cb8e89ded72d3ee80eacb5187b9996dd))
* use cluster definition in WAL service ([#17](https://github.com/cloudnative-pg/cnpg-i/issues/17)) ([4c0a1ac](https://github.com/cloudnative-pg/cnpg-i/commit/4c0a1ac46426d1c848d42f9f70b03cb1e07d5e22))
* use JSON patch instead of expecting the whole object ([#15](https://github.com/cloudnative-pg/cnpg-i/issues/15)) ([60e6b0f](https://github.com/cloudnative-pg/cnpg-i/commit/60e6b0f90cb4261817391d0c81e5f42685ff20da))
