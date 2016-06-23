# Roadmap

## PDEX : A CLI tool for PDexchange environments

Pdex is a great tool for Pdexchange development environments.

This tool comprises of the following flexibilities,i.e.,

- easy installation in windows, linux and mac
- scriptability
- scalability
- simple design
- simple interfacing
- stable

## PDExchange API Interaction from pdex cli is simple

*Authentication CLIs*

- pdex auth token create

- pdex auth token update
- pdex auth secret update


*Application CLIs*

- pdex apps list
- pdex apps app_id
- pdex apps app_id channels
- pdex apps app_id messages list
- pdex apps app_id messages latest
- pdex apps app_id messages msg_id

- pdex apps create
- pdex apps update app_id


*Channel CLIs*

- pdex channels channel_id messages list
- pdex channels channel_id messages msg_id
- pdex channels channel_id messages latest

- pdex channels add
- pdex channels channel_id command

- pdex channels channel_id remove


*Devicegroups CLIs*

- pdex devicegroups list
- pdex devicegroups deid_prefix
- pdex devicegroups deid_prefix devices list

- pdex devicegroups add
- pdex devicegroups deid_prefix new


*Devicegroups CLIs*

- pdex devices de_id list
- pdex devices de_id channels
- pdex pdex devices de_id commands
- devices de_id commands cmd_id

- pdex devices de_id message

*Self CLIs*

- pdex show me
- pdex delete me

- pdex setup
- pdex users

*Utils CLIs*

- pdex utils ping
- pdex utils version
- pdex utils access_token
- pdex utils secret_token de_id
- pdex utils app_token channel_id
- pdex utils setup
- pdex utils pdexadm
- pdex utils changelog

- pdex utils hmac
- pdex utils secret_token
