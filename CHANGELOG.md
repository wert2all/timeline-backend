# Changelog

## [1.12.1](https://github.com/wert2all/timeline-backend/compare/v1.12.0...v1.12.1) (2025-04-23)


### Bug Fixes

* **graph:** update timeline model in add-timeline resolver ([3610f8b](https://github.com/wert2all/timeline-backend/commit/3610f8b85e14e5357b065b383dba67a8e8fc7ba5))

## [1.12.0](https://github.com/wert2all/timeline-backend/compare/v1.11.1...v1.12.0) (2025-04-18)


### Features

* expose event ([d840ef1](https://github.com/wert2all/timeline-backend/commit/d840ef12e93e29b8021acc91facc2fed238077ca))
* expose timelineID with event type ([ee0a53b](https://github.com/wert2all/timeline-backend/commit/ee0a53b22d3f1ca5be53582e359cedbb68d1a6d9))
* **graph:** add account resolver for timeline ([7ffae59](https://github.com/wert2all/timeline-backend/commit/7ffae590fad4d4bd4b8c173df6a04d34308f52ae))
* **timeline:** add timeline resolver for TimelineEvent ([235124b](https://github.com/wert2all/timeline-backend/commit/235124b4deb4d8005d8d8ec0543afc5a3879f5c1))


### Bug Fixes

* fix static tests ([327023e](https://github.com/wert2all/timeline-backend/commit/327023e757b91d9e9c9619a84d73373a3b04858c))

## [1.11.1](https://github.com/wert2all/timeline-backend/compare/v1.11.0...v1.11.1) (2025-02-15)


### Bug Fixes

* fix saving avatarId ([4950dc1](https://github.com/wert2all/timeline-backend/commit/4950dc1d9ce3fdf53b769e0939536b0c62c23296))
* fix saving nil avatarID ([cba0e63](https://github.com/wert2all/timeline-backend/commit/cba0e630c7b3adb4d190714864ed4743313ecd2a))

## [1.11.0](https://github.com/wert2all/timeline-backend/compare/v1.10.0...v1.11.0) (2025-02-10)


### Features

* add about to account entity and GQL type ([54baa95](https://github.com/wert2all/timeline-backend/commit/54baa953f4c464cb835e8dbe18599d2b49b4a34f))

## [1.10.0](https://github.com/wert2all/timeline-backend/compare/v1.9.0...v1.10.0) (2025-02-07)


### Features

* add timeline query ([a02b6d1](https://github.com/wert2all/timeline-backend/commit/a02b6d1b7403bbce1117c7a73140ab4052b6bc38))


### Bug Fixes

* fix private events ([5b115f3](https://github.com/wert2all/timeline-backend/commit/5b115f3387418b2502efddc115149ab3e6ac6168))

## [1.9.0](https://github.com/wert2all/timeline-backend/compare/v1.8.0...v1.9.0) (2025-02-05)


### Features

* add code to auth GQL errors ([8e42030](https://github.com/wert2all/timeline-backend/commit/8e420308cac27cfbd12af1f8c52a746883698195))


### Bug Fixes

* accountId for expose timeline event is optional ([f1c7f29](https://github.com/wert2all/timeline-backend/commit/f1c7f2918adaebcd16a1999f9659f5df8c19a9e7))
* fix showing private events ([6795595](https://github.com/wert2all/timeline-backend/commit/67955956852c7f7bd699449cc136132adec1fdff))
* remove useless GQL operation ([4981245](https://github.com/wert2all/timeline-backend/commit/49812459a80effb76e6a99257dcc7c85cecfdc4f))
* return error on panic ([9a89fc0](https://github.com/wert2all/timeline-backend/commit/9a89fc06f60a445f232a00a071c7dbe213dd24de))

## [1.8.0](https://github.com/wert2all/timeline-backend/compare/v1.7.0...v1.8.0) (2025-01-27)


### Features

* add account GQL ([14a30a1](https://github.com/wert2all/timeline-backend/commit/14a30a17e287ccb51ddd1c621518ec8bcd100df7))
* expose avatarId of account ([6b9d60b](https://github.com/wert2all/timeline-backend/commit/6b9d60b6c3850e7c352b964716a430e3339697ab))
* saving account ([351140a](https://github.com/wert2all/timeline-backend/commit/351140a7f8925df0c08fa649ef3c35b5a49c0e4c))


### Bug Fixes

* fix removing event image ([d02fd40](https://github.com/wert2all/timeline-backend/commit/d02fd40b02060e9210459fa453fe3250a1ffe01c))
* fix static test ([8274114](https://github.com/wert2all/timeline-backend/commit/8274114b269b0dbd17d9907262425b0a0670e722))
* remove account avatar ([2a16042](https://github.com/wert2all/timeline-backend/commit/2a16042b60b45a74aa7ab61d873075e354e05570))

## [1.7.0](https://github.com/wert2all/timeline-backend/compare/v1.6.1...v1.7.0) (2025-01-16)


### Features

* expose event as cursor based pagination ([abd8002](https://github.com/wert2all/timeline-backend/commit/abd8002b57f5a1425dd1be9d0feea6a0982af5ea))


### Bug Fixes

* disable introspaction for production ([7160a96](https://github.com/wert2all/timeline-backend/commit/7160a96d7fd25f28b9b41b899afecd36a649ba17))
* token should be base64 encoded ([aae46c1](https://github.com/wert2all/timeline-backend/commit/aae46c12e9546c74223037f5674adc1f3c3df346))

## [1.6.1](https://github.com/wert2all/timeline-backend/compare/v1.6.0...v1.6.1) (2025-01-13)


### Bug Fixes

* fix go mod ([13dffb5](https://github.com/wert2all/timeline-backend/commit/13dffb5350c846d29ab4b361e3b10e6f1526ec09))
* fix gql server handler ([745dcba](https://github.com/wert2all/timeline-backend/commit/745dcba955c913673007fa5b650198136287033d))
* fix saving date ([7adce94](https://github.com/wert2all/timeline-backend/commit/7adce944e041d248492fa854af15641616b8120c))
* fix showing showTime after update ([4408ccb](https://github.com/wert2all/timeline-backend/commit/4408ccbe5cb252226215df8b6afbfa9b0c87809b))
* fix url value after update event ([2e9db3b](https://github.com/wert2all/timeline-backend/commit/2e9db3bafd320b8ac97a08f10209d36b4089c1fd))

## [1.6.0](https://github.com/wert2all/timeline-backend/compare/v1.5.0...v1.6.0) (2024-12-31)


### Features

* add previewly GQL scheme ([64e1433](https://github.com/wert2all/timeline-backend/commit/64e14332ea28cd3cbd4b30900d102a4bd42d771f))
* add previewly token to account ([95f427d](https://github.com/wert2all/timeline-backend/commit/95f427dedb145931f3f646408e171fb38f5674dc))
* expose account previewly token ([92c3a04](https://github.com/wert2all/timeline-backend/commit/92c3a044a1809fe2a44a8df3706d1638bef4819d))
* expose previewly image ID ([640c334](https://github.com/wert2all/timeline-backend/commit/640c33416eabea437323e65fc1b7911cb47ec019))
* expose token from previewly service ([df1ac02](https://github.com/wert2all/timeline-backend/commit/df1ac028f673ce5fdcf6521f6297c3a5512fecf1))
* saving previewly image id ([d3f79cb](https://github.com/wert2all/timeline-backend/commit/d3f79cb3446ce20f7cc85b398f96e1e2cd04991b))


### Bug Fixes

* docker container should not restart always ([6a52825](https://github.com/wert2all/timeline-backend/commit/6a528257894e1d1180ad1655306b7f2d25f5e5d8))
* fix adding timeline ([bf13ce8](https://github.com/wert2all/timeline-backend/commit/bf13ce883ccab2d7737e6d93d35b007e1badd7bc))
* fix deleting event ([4de9f17](https://github.com/wert2all/timeline-backend/commit/4de9f17d02af11b573c3621c5974f43dd9d936cc))
* fix error shadow ([313742c](https://github.com/wert2all/timeline-backend/commit/313742ce6ae0017a679fe69d32b10223fb4b246b))
* fix saving image ID on adding event ([bbd9043](https://github.com/wert2all/timeline-backend/commit/bbd9043701a77d588d16910466c08f5ee476d2f3))
* fix static test ([de12a12](https://github.com/wert2all/timeline-backend/commit/de12a12d4707c2bbe04ebab4be60032894afc93a))
* fix static test ([ab10517](https://github.com/wert2all/timeline-backend/commit/ab10517adaba6c2cb3aca2d1cdc9c482c900bb16))
* remove auth middleware at all ([e2977a5](https://github.com/wert2all/timeline-backend/commit/e2977a5c0797f05d709b329463f2f19ee11d389f))

## [1.5.0](https://github.com/wert2all/timeline-backend/compare/v1.4.1...v1.5.0) (2024-12-08)


### Features

* (db) create settings entity ([cc59338](https://github.com/wert2all/timeline-backend/commit/cc59338be6041b3139edb8d2529f95ad0b95bda1))
* (db) create settings model ([75134ef](https://github.com/wert2all/timeline-backend/commit/75134efe9e9802f187cfff6b7aeaec9590bee18d))
* (gql) expose account settings ([cf062bf](https://github.com/wert2all/timeline-backend/commit/cf062bf65258a932ec1014be81b0919e5a413dd0))
* (gql) expose account settings ([70cdb55](https://github.com/wert2all/timeline-backend/commit/70cdb55e15c841471323d54055baf0e044e40859))
* add GQL query for expose account timelines ([0262dd1](https://github.com/wert2all/timeline-backend/commit/0262dd1a100df8a530c28b5ddf3a4c431da4bab5))
* add interfaces for saveSettings mutation ([1b69441](https://github.com/wert2all/timeline-backend/commit/1b694419507a1a1f9bd1dc92e0c815b179af727d))
* add save setting GQL mutation ([4e9d0ff](https://github.com/wert2all/timeline-backend/commit/4e9d0ff89552fb443d5b6a081beef72c4309c006))
* add token to context ([9e7456c](https://github.com/wert2all/timeline-backend/commit/9e7456c8e9584db66ed46dae695d67095b5d6e83))
* saving account setings ([a75bf80](https://github.com/wert2all/timeline-backend/commit/a75bf80de7f3e8119f87516b9d50a66d4aae60b4))
* validate input of saveSettings method ([3b8f4e8](https://github.com/wert2all/timeline-backend/commit/3b8f4e8bb0c3d7367204608088e75b1232c977c4))


### Bug Fixes

* fix assign event to timeline ([bf448d4](https://github.com/wert2all/timeline-backend/commit/bf448d4527ccb0b9996816f4630eb234b5fb6279))
* rename account type on authorization ([5e65d8e](https://github.com/wert2all/timeline-backend/commit/5e65d8ef2d562d7cde72fd89ca26363e87e17ca5))
* upsert account setting ([f5758b8](https://github.com/wert2all/timeline-backend/commit/f5758b880d2e1bcb6ac36979092d302a602ca151))

## [1.4.1](https://github.com/wert2all/timeline-backend/compare/v1.4.0...v1.4.1) (2024-11-07)


### Bug Fixes

* add id to account ([3d3bad2](https://github.com/wert2all/timeline-backend/commit/3d3bad2cdcd82d6bb8c9bb5f2ca501bdcd2e0334))
* make user email unique and non nullable ([156be3b](https://github.com/wert2all/timeline-backend/commit/156be3bf0f1f79b7c2e115cfac926a1b273d908d))

## [1.4.0](https://github.com/wert2all/timeline-backend/compare/v1.3.0...v1.4.0) (2024-11-06)


### Features

* add account ent entity ([a9db9b5](https://github.com/wert2all/timeline-backend/commit/a9db9b5a2db94f39ffcab38944ea5a61be6ecb9b))
* add account repository ([87791d5](https://github.com/wert2all/timeline-backend/commit/87791d5d7230dedb546d11d3584374f308c0404d))
* create account after first login ([67272fc](https://github.com/wert2all/timeline-backend/commit/67272fc679916a39122f754912a55dea6d6c6ca8))
* expose user accounts ([2488821](https://github.com/wert2all/timeline-backend/commit/248882149d562acae618fd03c406480dcc2a9e2e))

## [1.3.0](https://github.com/wert2all/timeline-backend/compare/v1.2.0...v1.3.0) (2024-10-31)


### Features

* saving existing event ([446dc4d](https://github.com/wert2all/timeline-backend/commit/446dc4ddf7db7f49e354f111a04096260d7affa7))


### Bug Fixes

* fix DI of base validator ([f421cd1](https://github.com/wert2all/timeline-backend/commit/f421cd1b2c8d9afb2b0f26356edb73db1d500fa8))
* fix static test ([ca69f4e](https://github.com/wert2all/timeline-backend/commit/ca69f4e3193fdadde0f51d2b68339ac1260f34ed))
* remove usless dependencues from add event validator ([29b6a8b](https://github.com/wert2all/timeline-backend/commit/29b6a8b694d693186711b5bba7fa4149a08b28d6))

## [1.2.0](https://github.com/wert2all/timeline-backend/compare/v1.1.8...v1.2.0) (2024-10-20)


### Features

* add custom logger ([4af6fdd](https://github.com/wert2all/timeline-backend/commit/4af6fdd0f7e6fdb26899ecf649cfcbde037c829b))
* add utils D method ([b6012c4](https://github.com/wert2all/timeline-backend/commit/b6012c4260f7fa32d17dc7291aa607298ad98077))


### Bug Fixes

* change postgres port for docker compose ([009ac31](https://github.com/wert2all/timeline-backend/commit/009ac318408c905c385b429739fef7fbc45c8fa1))
* fix container ([67428f1](https://github.com/wert2all/timeline-backend/commit/67428f11522efc14dcb5f783c1315eb91afbf658))
* fix extract user ID from context ([e1b3903](https://github.com/wert2all/timeline-backend/commit/e1b39034ff882ec9a0d7e82dc73c09c7e4e497b4))
* fix generated code ([cdcbeee](https://github.com/wert2all/timeline-backend/commit/cdcbeeed6033a9b1e8006579016b37a4c8f51c3e))
* fix golang-lint action ([a0ec4e6](https://github.com/wert2all/timeline-backend/commit/a0ec4e6badc975ba06c5b3021a044cc7e0909db7))
* fix GQL routing ([83c232c](https://github.com/wert2all/timeline-backend/commit/83c232c3e42b8dd6a71b6068ae2606d78fbddc7e))
* fix parsing command prompt string ([b54e2cf](https://github.com/wert2all/timeline-backend/commit/b54e2cf9c03bdba2023141089521726616ffe497))
* fix sql driver ([a3e443f](https://github.com/wert2all/timeline-backend/commit/a3e443f941458943d9cde25d664a80c3a2eb778e))
* remove useless code ([89a6d1e](https://github.com/wert2all/timeline-backend/commit/89a6d1e2fba7d99f3fad189cd2199c1244d1989b))

## [1.1.8](https://github.com/wert2all/timeline-backend/compare/v1.1.7...v1.1.8) (2024-10-16)


### Bug Fixes

* fix go version ([6ce8bf8](https://github.com/wert2all/timeline-backend/commit/6ce8bf843fa61a19e787091e0345c19db012c4d3))
* fix nullable email and name of google user ([6c909f4](https://github.com/wert2all/timeline-backend/commit/6c909f40c6eb4e79fab04a8c07f426c0beabce27))
* fix version of chi ([83496ac](https://github.com/wert2all/timeline-backend/commit/83496ace1ac0b7995cc69a78b6da04f233d94e3b))

## [1.1.7](https://github.com/wert2all/timeline-backend/compare/v1.1.6...v1.1.7) (2024-05-25)


### Bug Fixes

* fix docker again ([7c65867](https://github.com/wert2all/timeline-backend/commit/7c6586745089527dc709c8b116adeb881a46f73d))

## [1.1.6](https://github.com/wert2all/timeline-backend/compare/v1.1.5...v1.1.6) (2024-05-25)


### Bug Fixes

* fix dockerfile ([a9818ef](https://github.com/wert2all/timeline-backend/commit/a9818effe8d856be777589fc5dd4e60786858c38))

## [1.1.5](https://github.com/wert2all/timeline-backend/compare/v1.1.4...v1.1.5) (2024-05-25)


### Bug Fixes

* fix docker ([fb8c686](https://github.com/wert2all/timeline-backend/commit/fb8c6864d93254f8c67b10d897d607ee27e25921))

## [1.1.4](https://github.com/wert2all/timeline-backend/compare/v1.1.3...v1.1.4) (2024-05-24)


### Bug Fixes

* add dump for application config ([040ac09](https://github.com/wert2all/timeline-backend/commit/040ac09ec1daab49d848e798de040185dc7ab22e))
* fix docker compose env ([3377406](https://github.com/wert2all/timeline-backend/commit/3377406eff0504782d2708687cea5f9c2ddae715))
* remove config ([ca47701](https://github.com/wert2all/timeline-backend/commit/ca47701aed47061d6db5447ee09d3fc76af44b70))

## [1.1.3](https://github.com/wert2all/timeline-backend/compare/v1.1.2...v1.1.3) (2024-05-23)


### Bug Fixes

* fix postgres driver ([6c9b8f1](https://github.com/wert2all/timeline-backend/commit/6c9b8f1e11f9565767821e7566af15aa66b3125b))
* fix static test ([90912bf](https://github.com/wert2all/timeline-backend/commit/90912bf45f5d3995a6a4a8dad412b1f322a10239))
* move creating client etc to di ([442e255](https://github.com/wert2all/timeline-backend/commit/442e255e9484c0ac07571bd89c82338a59d3f019))
* remove CLose() from Application ([e19c648](https://github.com/wert2all/timeline-backend/commit/e19c6485cb73ac3e05775dbfbef06e319737345c))
* remove useless function ([0df0191](https://github.com/wert2all/timeline-backend/commit/0df0191dd2f58cb02a80fcdf3d9cef2ff98c4727))

## [1.1.2](https://github.com/wert2all/timeline-backend/compare/v1.1.1...v1.1.2) (2024-05-22)


### Bug Fixes

* fix docker image build ([21ff985](https://github.com/wert2all/timeline-backend/commit/21ff985d8edde3e8e5409ecbc046e2278257abbc))

## [1.1.1](https://github.com/wert2all/timeline-backend/compare/v1.1.0...v1.1.1) (2024-05-22)


### Bug Fixes

* fix docker image name ([4abbc88](https://github.com/wert2all/timeline-backend/commit/4abbc8842eccc0491046d1045f9d6e38800dfd4a))

## [1.1.0](https://github.com/wert2all/timeline-backend/compare/v1.0.0...v1.1.0) (2024-05-22)


### Features

* add chi Recoverer middleware ([3b0d73f](https://github.com/wert2all/timeline-backend/commit/3b0d73fbe412b3e015473c99db5b1fe9ca97c841))
* add dumper ([a6d95b5](https://github.com/wert2all/timeline-backend/commit/a6d95b58de519629f745790e4d269270c5389796))
* add tags to docker image ([fd670a3](https://github.com/wert2all/timeline-backend/commit/fd670a334dd0e279bc55f4af9f7e56be0eb11ad2))
* refactor config to yaml ([061f49c](https://github.com/wert2all/timeline-backend/commit/061f49cad42ea5000c263daa05dfe58bdac913ed))


### Bug Fixes

* add registry to docker image ([b48afab](https://github.com/wert2all/timeline-backend/commit/b48afabfff83b58ba782bc6063a7991b5f8ceca1))
* fix build ([d8697e1](https://github.com/wert2all/timeline-backend/commit/d8697e1ee08885f51799705a3d8161be299416a9))
* fix go.sum ([322bbe3](https://github.com/wert2all/timeline-backend/commit/322bbe3665bef5ad7cf202c35accbed30fcc4d91))
* remove port from application ([6214a00](https://github.com/wert2all/timeline-backend/commit/6214a00418c26fad7ed7e046f031156b6a69a4ed))

## 1.0.0 (2024-05-20)


### Features

* add .env.dist ([2763b2b](https://github.com/wert2all/timeline-backend/commit/2763b2b8254956ec8f4daa5a89f35c89b41af6e6))
* add authorize method to GQL Mutation ([b666e36](https://github.com/wert2all/timeline-backend/commit/b666e36c0cf445c9b322f777e2e88ba5e1196983))
* add context ([a0a16bc](https://github.com/wert2all/timeline-backend/commit/a0a16bc645f59d8746dad05d5c16b658a016e20f))
* add CORS as middleware ([90acebd](https://github.com/wert2all/timeline-backend/commit/90acebde5a550a87708b2b2a0219c6801affb654))
* add delete event mutation ([b29c5f2](https://github.com/wert2all/timeline-backend/commit/b29c5f2bc351afac322ae4e0454a5cd0b8d1c37e))
* add docker compose support ([63e3eb5](https://github.com/wert2all/timeline-backend/commit/63e3eb5b8968fe24211931f78d0d5f03aeb6970e))
* add dotenv files support ([84ff1d2](https://github.com/wert2all/timeline-backend/commit/84ff1d2aba9777b52d54eb17b10c6d7e7bfafc37))
* add ent support ([e56a816](https://github.com/wert2all/timeline-backend/commit/e56a816b6360b2f0418240bff6f3331f60f68453))
* add gaper binary file to gitignore ([cf76247](https://github.com/wert2all/timeline-backend/commit/cf762479ecd4301708c526be9eca0e32b1b9276d))
* add getting timeline events ([f4a940d](https://github.com/wert2all/timeline-backend/commit/f4a940d0d2a64d9144885546f37b1f83c5a132d6))
* add google api client ([4b32c06](https://github.com/wert2all/timeline-backend/commit/4b32c0619763bab7808086b22fafee8a949e93c2))
* add isNew property to user type ([1e02395](https://github.com/wert2all/timeline-backend/commit/1e0239509491a17a0980e00bf6ca262846dd5371))
* add method for addint timeline events ([6543138](https://github.com/wert2all/timeline-backend/commit/65431388fb2771e8f0b88cdeff55b0f3a02b85b7))
* add ordering to list of events ([6157264](https://github.com/wert2all/timeline-backend/commit/6157264e0c9d9b620c83ec2e2770dc0f80ff10b0))
* add postgres settings to config ([3d1f0c9](https://github.com/wert2all/timeline-backend/commit/3d1f0c95c311776882c07ad8029121bc440c459c))
* add README.md ([238cbd2](https://github.com/wert2all/timeline-backend/commit/238cbd22594cc796b50368fa8ee0de0fbc67381d))
* add relation to timeline and user ([74610f7](https://github.com/wert2all/timeline-backend/commit/74610f7fc8d240fe31f6c1f17cfd3c72a414543f))
* add sentry support ([a6ead27](https://github.com/wert2all/timeline-backend/commit/a6ead275ad776f0e00131d208e31067b04507e78))
* add show time to entity type ([9fc0c20](https://github.com/wert2all/timeline-backend/commit/9fc0c203b0524a713350a7ff0727999376c5de68))
* add simple graphql server ([7e2b7bb](https://github.com/wert2all/timeline-backend/commit/7e2b7bb19c9c16714a32cb7083297a96d71b89ac))
* add tag entity ([ab51349](https://github.com/wert2all/timeline-backend/commit/ab51349040f38b9ef8d9d72d122e7650439f4955))
* add tag repository to sl and to model ([fd58290](https://github.com/wert2all/timeline-backend/commit/fd58290f8c7abf24ecc78872480d7e480f1a7462))
* add tags for added event ([fb4f31f](https://github.com/wert2all/timeline-backend/commit/fb4f31f2c730b2f57f7b6f5cd2a4c9b4d8afc590))
* add tags to add event GQL ([e99701a](https://github.com/wert2all/timeline-backend/commit/e99701a1a4db3246e6c32eb5760f36847d73f4b3))
* add tags to valid GQL input ([4d483f9](https://github.com/wert2all/timeline-backend/commit/4d483f991ff58e087c11376f4ce0249b538114e6))
* add timeline and event entities ([6dffb41](https://github.com/wert2all/timeline-backend/commit/6dffb41f1f4f7a247daeefbc769bd5155dd45a82))
* add timeline list to user GQL ([fe50d64](https://github.com/wert2all/timeline-backend/commit/fe50d64d5cdf07e8d3ac4b09b17c6159a1903182))
* add timeline model ([3ccd5eb](https://github.com/wert2all/timeline-backend/commit/3ccd5eb1f45ab724ef98c8d6229519da2ac370a3))
* add timeline mutation ([08aaa0f](https://github.com/wert2all/timeline-backend/commit/08aaa0f1bc5e76cfe6fc77f482c9a603acb181cf))
* add type to timeline event entity ([2b9c44d](https://github.com/wert2all/timeline-backend/commit/2b9c44d4a29e1c37cf0d332708550b305e942f6b))
* add url by GQL ([ac960ed](https://github.com/wert2all/timeline-backend/commit/ac960eddfc2c34f1d58c4feee2fa905a1786a478))
* add url field to db ([069d60d](https://github.com/wert2all/timeline-backend/commit/069d60df1db21212c69d9f42ba3a982d6c29f9da))
* add vscode support ([96a77e5](https://github.com/wert2all/timeline-backend/commit/96a77e587a5b03287d33278b7cd0a897c6da097d))
* converting timelines ([0ea8e0c](https://github.com/wert2all/timeline-backend/commit/0ea8e0c4b6d136e86366d0f93a9482bcf97bce6a))
* create hello world ([7a66f29](https://github.com/wert2all/timeline-backend/commit/7a66f29867b018c437ee49b58819e6b391daca83))
* create user on request to API ([78f42fe](https://github.com/wert2all/timeline-backend/commit/78f42fe5840843e248ec5d309896ef66d3b5431d))
* etract token from request ([48ae7ee](https://github.com/wert2all/timeline-backend/commit/48ae7ee50fa8030c033cbfdadcff9bb8f046b70d))
* expose tags for every event entity ([5dc4588](https://github.com/wert2all/timeline-backend/commit/5dc45885be81ad219da4af9463fd9463835d5f6a))
* extract google user from payload ([642a048](https://github.com/wert2all/timeline-backend/commit/642a0480c77fc02e5ad5427d398b39091859c9e1))
* return user by authorization mutation ([33ae550](https://github.com/wert2all/timeline-backend/commit/33ae550a244a4267a3f07862420622fcd4f7498b))
* save description ([047691f](https://github.com/wert2all/timeline-backend/commit/047691f05a5d7c8deb51cd84af160b8c004d3ba7))
* saving title ([9705173](https://github.com/wert2all/timeline-backend/commit/970517354b32b1332cc172cf37e29dadec9d8943))
* upset tag ([214d3bf](https://github.com/wert2all/timeline-backend/commit/214d3bf1e5252741f7d7953c862cda368798109e))


### Bug Fixes

* add avatar to GQL User type ([19d4937](https://github.com/wert2all/timeline-backend/commit/19d4937b40834cc3824e5341c6884bfac9ef90ee))
* add error to every methods of interfaces ([2cdc0e0](https://github.com/wert2all/timeline-backend/commit/2cdc0e04882c6fa536b36b83c4b7a15d91653469))
* add validation to add timeline method ([eea668e](https://github.com/wert2all/timeline-backend/commit/eea668eae97d248ce4e1f4aea6cd5fd8179e1c14))
* disable loading vscode config on nvim with neoconfig plugin ([16db420](https://github.com/wert2all/timeline-backend/commit/16db420cc76b7fa566b6c21b77d5edc8fbae3df7))
* fix add timeline GQL method ([e53609c](https://github.com/wert2all/timeline-backend/commit/e53609c61ba34dd07964b65659708ad0cddee585))
* fix adding empty tags ([220d0b9](https://github.com/wert2all/timeline-backend/commit/220d0b9bcaeb2d5af93410c336784b484f95dbbf))
* fix adding tag ([5acfbeb](https://github.com/wert2all/timeline-backend/commit/5acfbeb794f96ce53e0047c6401a38c5b877a2ed))
* fix application ([171b5ac](https://github.com/wert2all/timeline-backend/commit/171b5ac49d5ce71aca61e3ef507aa3615a2ea6d4))
* fix application ([4eee788](https://github.com/wert2all/timeline-backend/commit/4eee788017ba55b5736120acfd2ed5a50c844d6f))
* fix authorization ([350e23c](https://github.com/wert2all/timeline-backend/commit/350e23cd0d70ad1a7f49bc198b225fbd939d14bc))
* fix authorize middleware to use on post requests only ([eb0d043](https://github.com/wert2all/timeline-backend/commit/eb0d043a73baa6eb3881abcc13859859f5a7f9f3))
* fix authorize midleware to use on post requests only ([4347ba8](https://github.com/wert2all/timeline-backend/commit/4347ba8371b5f514fe439fec8bee1e5075954771))
* fix build ([5c10aa9](https://github.com/wert2all/timeline-backend/commit/5c10aa9b896964e0651f17bfbd55d6c382c9203c))
* fix cicd ([c8e7771](https://github.com/wert2all/timeline-backend/commit/c8e7771e59839a42c987e71ac6f20934c7f7bf30))
* fix cicd build ([27cd23b](https://github.com/wert2all/timeline-backend/commit/27cd23ba219a31180ea7efd3420499f719c4b659))
* fix docker compose ([84695c9](https://github.com/wert2all/timeline-backend/commit/84695c9036fba2ba681e92eeb8a5c20b2e196c1c))
* fix empty tags for add event GQL ([6325776](https://github.com/wert2all/timeline-backend/commit/6325776655b8d3c29d6d4d9e8632a4f522efb479))
* fix expose empty tags ([f0d8ae0](https://github.com/wert2all/timeline-backend/commit/f0d8ae07c33743ece383591effd275d688854ace))
* fix expose tags ([da66098](https://github.com/wert2all/timeline-backend/commit/da66098ab6092562e0d93c64c91e587691af6c65))
* fix gql generation ([06df58c](https://github.com/wert2all/timeline-backend/commit/06df58ced95cb1dfc6616605dc92a18266e91995))
* fix mod ([6f8dbe6](https://github.com/wert2all/timeline-backend/commit/6f8dbe61d4f2134ad119cf8a6578c733a9eeeac9))
* fix reading dotenv variables ([353dd95](https://github.com/wert2all/timeline-backend/commit/353dd9596d07a7dc13619aaf65ba9a3629406ec6))
* fix showTime input argument ([02ca3c9](https://github.com/wert2all/timeline-backend/commit/02ca3c9d21263c24a9ddbbdb8c2428a45451b15d))
* fix types of serviceLocator ([9a37a11](https://github.com/wert2all/timeline-backend/commit/9a37a11cde4b31d70a949e3056f9823f229ee34a))
* fix URL validation ([2a51dbf](https://github.com/wert2all/timeline-backend/commit/2a51dbf2bd443f9b911e4ed3d2bec24ab03c5c64))
* fix vscode settings ([35761c2](https://github.com/wert2all/timeline-backend/commit/35761c2bc8340e02a0918d41b0ea1a693221cb5b))
* make user repository implementation as private ([c9c28a5](https://github.com/wert2all/timeline-backend/commit/c9c28a587dd368837d926235d9cab4a543b4a7d6))
* remove id from GQL types ([1255bf9](https://github.com/wert2all/timeline-backend/commit/1255bf975b1a8ae4a02a056b1a75aca61b316107))
* remove useless resolvers from DI ([dde5554](https://github.com/wert2all/timeline-backend/commit/dde55540ba408da713263036656b32c9673cce40))
* sanitize text inputs ([c255ba7](https://github.com/wert2all/timeline-backend/commit/c255ba773ebe780fcc1953f3b4891aa937aca428))
* update time when user was authorized ([a887f6f](https://github.com/wert2all/timeline-backend/commit/a887f6f9f2bff7c820d924930a3337e48bcbfcf9))
