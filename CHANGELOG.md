# Changelog

## 0.1.0 (2024-12-24)


### Features

* add `deregister` capabilities ([#75](https://github.com/cloudnative-pg/cnpg-i/issues/75)) ([9ea12e7](https://github.com/cloudnative-pg/cnpg-i/commit/9ea12e76a6ee41bc625cd81d9e66517b092299aa))
* add backup service ([#19](https://github.com/cloudnative-pg/cnpg-i/issues/19)) ([92ca0ea](https://github.com/cloudnative-pg/cnpg-i/commit/92ca0eaa10ad0dc88b7458b56331092453e9035f))
* add lifecycle hooks ([#20](https://github.com/cloudnative-pg/cnpg-i/issues/20)) ([5869c64](https://github.com/cloudnative-pg/cnpg-i/commit/5869c644788b8d4d3de0b5305bd0fc3c7f1be17b))
* add operator service type in the plugin metadata ([#6](https://github.com/cloudnative-pg/cnpg-i/issues/6)) ([7fcdddb](https://github.com/cloudnative-pg/cnpg-i/commit/7fcdddbaf7f46c89affcf3871ddd9fb51e691adc))
* add plugin metadata ([#7](https://github.com/cloudnative-pg/cnpg-i/issues/7)) ([e16cc54](https://github.com/cloudnative-pg/cnpg-i/commit/e16cc543b9bb875b02d10b4cfa015ebd3050ea22)), closes [#4](https://github.com/cloudnative-pg/cnpg-i/issues/4)
* add pre and post reconcile hooks ([#21](https://github.com/cloudnative-pg/cnpg-i/issues/21)) ([af27643](https://github.com/cloudnative-pg/cnpg-i/commit/af2764325710cd98881e211810d2a01f61b62d43))
* allow plugins to set the status ([#67](https://github.com/cloudnative-pg/cnpg-i/issues/67)) ([cf1ab8f](https://github.com/cloudnative-pg/cnpg-i/commit/cf1ab8fab2a05e7e28cb7ff1aa575a447d9e6634))
* operator reconciler specification ([#3](https://github.com/cloudnative-pg/cnpg-i/issues/3)) ([0b01870](https://github.com/cloudnative-pg/cnpg-i/commit/0b018708a768a8f4d59e99716a782e44cceb7ff2))
* pluggable dynamic controller ([#25](https://github.com/cloudnative-pg/cnpg-i/issues/25)) ([7d938f6](https://github.com/cloudnative-pg/cnpg-i/commit/7d938f60f998af1d336a669a2497d1cf68b9d582))
* postgres interface ([#42](https://github.com/cloudnative-pg/cnpg-i/issues/42)) ([b8e17e4](https://github.com/cloudnative-pg/cnpg-i/commit/b8e17e42e05c1dc5b640a1bbf2466b1fc193d267))
* remove MutatePod method ([#26](https://github.com/cloudnative-pg/cnpg-i/issues/26)) ([9afc581](https://github.com/cloudnative-pg/cnpg-i/commit/9afc58113460f207ca4fb83b6ba875caef1abb64))
* restore job entrypoint and capability ([#110](https://github.com/cloudnative-pg/cnpg-i/issues/110)) ([c704f46](https://github.com/cloudnative-pg/cnpg-i/commit/c704f46c20e0f2d2b4c1c1673ce55b25a7b972f9))
* service for management of WAL files ([#2](https://github.com/cloudnative-pg/cnpg-i/issues/2)) ([963073d](https://github.com/cloudnative-pg/cnpg-i/commit/963073d61df217fd6c2cd9f4ca77b4a57e623c04)), closes [#1](https://github.com/cloudnative-pg/cnpg-i/issues/1)


### Bug Fixes

* add cluster definition to MutatePod RPC call ([#10](https://github.com/cloudnative-pg/cnpg-i/issues/10)) ([2492dba](https://github.com/cloudnative-pg/cnpg-i/commit/2492dba8ebbd98f1da02a9e59d05e5a76beac8e6))
* **deps:** update all non-major go dependencies ([1ee4c59](https://github.com/cloudnative-pg/cnpg-i/commit/1ee4c59422aa56005b2c3f9394fc817c76275511))
* **deps:** update all non-major go dependencies ([caf6b1b](https://github.com/cloudnative-pg/cnpg-i/commit/caf6b1b23daa7a733ad2a6c321ab097660bbdd94))
* **deps:** update module google.golang.org/grpc to v1.65.0 ([94bb720](https://github.com/cloudnative-pg/cnpg-i/commit/94bb720bd7a555d5176f66f2793d1b562b550ca0))
* **deps:** update module google.golang.org/grpc to v1.66.0 ([f231bc4](https://github.com/cloudnative-pg/cnpg-i/commit/f231bc4189df9dcf604320b59cc34e9d553a7ec1))
* **deps:** update module google.golang.org/grpc to v1.66.1 ([956b196](https://github.com/cloudnative-pg/cnpg-i/commit/956b196c4d303a2a6ee13581de35c5617cd8ec92))
* **deps:** update module google.golang.org/grpc to v1.66.2 ([e35a41f](https://github.com/cloudnative-pg/cnpg-i/commit/e35a41f25a3cbf4e0d878fa14eb52015c3cd7915))
* **deps:** update module google.golang.org/grpc to v1.67.0 ([#98](https://github.com/cloudnative-pg/cnpg-i/issues/98)) ([6926a08](https://github.com/cloudnative-pg/cnpg-i/commit/6926a08b23857be1c9b6e125fb9d83758566472c))
* **deps:** update module google.golang.org/grpc to v1.67.1 ([7e24b2e](https://github.com/cloudnative-pg/cnpg-i/commit/7e24b2eccd507275f994b458189843d6d52a0f4c))
* **deps:** update module google.golang.org/grpc to v1.68.0 ([#121](https://github.com/cloudnative-pg/cnpg-i/issues/121)) ([d4d5b6a](https://github.com/cloudnative-pg/cnpg-i/commit/d4d5b6a7d6eda0028b087ade2cbd24c7c7c72adb))
* **deps:** update module google.golang.org/grpc to v1.68.1 ([#130](https://github.com/cloudnative-pg/cnpg-i/issues/130)) ([20a47b2](https://github.com/cloudnative-pg/cnpg-i/commit/20a47b2d919b1bfaf9f759a4d6d7e857fd6ab960))
* **deps:** update module google.golang.org/grpc to v1.69.0 ([#135](https://github.com/cloudnative-pg/cnpg-i/issues/135)) ([a34651f](https://github.com/cloudnative-pg/cnpg-i/commit/a34651fddaa6902c95f4cdc33841eae67cb9b55f))
* **deps:** update module google.golang.org/grpc to v1.69.2 ([#140](https://github.com/cloudnative-pg/cnpg-i/issues/140)) ([cbc4287](https://github.com/cloudnative-pg/cnpg-i/commit/cbc4287931eec44707a228028ead90ac53d9a989))
* **deps:** update module google.golang.org/protobuf to v1.35.1 ([b209f8e](https://github.com/cloudnative-pg/cnpg-i/commit/b209f8e465716379308371ae5b48371159a50a97))
* **deps:** update module google.golang.org/protobuf to v1.35.2 ([#123](https://github.com/cloudnative-pg/cnpg-i/issues/123)) ([04f8945](https://github.com/cloudnative-pg/cnpg-i/commit/04f894529500f5721c10067a8e4459186a6a9110))
* **deps:** update module google.golang.org/protobuf to v1.36.0 ([#139](https://github.com/cloudnative-pg/cnpg-i/issues/139)) ([e526410](https://github.com/cloudnative-pg/cnpg-i/commit/e52641093a69541bafd197cc73f795c6308817c5))
* **deps:** update module google.golang.org/protobuf to v1.36.1 ([#141](https://github.com/cloudnative-pg/cnpg-i/issues/141)) ([bc221c3](https://github.com/cloudnative-pg/cnpg-i/commit/bc221c3954e3e4e7147a92c03809381364fc66ab))
* package names ([#30](https://github.com/cloudnative-pg/cnpg-i/issues/30)) ([b0b3107](https://github.com/cloudnative-pg/cnpg-i/commit/b0b310788fa1c097e139e31d661f124be91b94b7))
* reconciler package path ([#27](https://github.com/cloudnative-pg/cnpg-i/issues/27)) ([3624915](https://github.com/cloudnative-pg/cnpg-i/commit/362491594dde7a5360b24a0fd14c7e02715cab09))
* regenerate stub & skeleton ([#11](https://github.com/cloudnative-pg/cnpg-i/issues/11)) ([bc5f2ca](https://github.com/cloudnative-pg/cnpg-i/commit/bc5f2cab5e5949cc4e6dd0d2200613aea7330f20))
* remove maintainers from plugin metadata ([#12](https://github.com/cloudnative-pg/cnpg-i/issues/12)) ([a18f8f5](https://github.com/cloudnative-pg/cnpg-i/commit/a18f8f50cf36d1928a5d41ebeba192228cf6d2ea))
* remove stale constants ([#28](https://github.com/cloudnative-pg/cnpg-i/issues/28)) ([4adc035](https://github.com/cloudnative-pg/cnpg-i/commit/4adc03536d626577c919718d7d02bef89fbe366a))
* service capability order ([#14](https://github.com/cloudnative-pg/cnpg-i/issues/14)) ([bc73ed3](https://github.com/cloudnative-pg/cnpg-i/commit/bc73ed38c383c5e2b90c95ccb67a724e8354d5e7))
* use bytes for JSON serialization instead of strings ([#13](https://github.com/cloudnative-pg/cnpg-i/issues/13)) ([e99be6d](https://github.com/cloudnative-pg/cnpg-i/commit/e99be6d0cb8e89ded72d3ee80eacb5187b9996dd))
* use cluster definition in WAL service ([#17](https://github.com/cloudnative-pg/cnpg-i/issues/17)) ([4c0a1ac](https://github.com/cloudnative-pg/cnpg-i/commit/4c0a1ac46426d1c848d42f9f70b03cb1e07d5e22))
* use JSON patch instead of expecting the whole object ([#15](https://github.com/cloudnative-pg/cnpg-i/issues/15)) ([60e6b0f](https://github.com/cloudnative-pg/cnpg-i/commit/60e6b0f90cb4261817391d0c81e5f42685ff20da))
* workflow name ([#71](https://github.com/cloudnative-pg/cnpg-i/issues/71)) ([3b53627](https://github.com/cloudnative-pg/cnpg-i/commit/3b53627724590185381246c0e256ecf1de55c99d))