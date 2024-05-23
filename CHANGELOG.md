# Changelog

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
