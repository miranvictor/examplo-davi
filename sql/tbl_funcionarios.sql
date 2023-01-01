create table public.funcionarios
(
    id                    integer default nextval('servidores_id_seq'::regclass) not null
        constraint servidores_pkey
            primary key,
    cd_uj                 integer,
    cd_upag               varchar,
    ds_nome               varchar,
    cd_cpf                varchar,
    cd_matricula          varchar,
    id_regima             integer,
    ds_cargo              varchar,
    id_natureza_cargo     integer,
    dt_exercicio          varchar,
    dt_aposentadoria      varchar,
    dt_exclusao           varchar,
    nr_jornada            integer,
    cd_categoria_situacao integer,
    dt_mes_ano_pagamento  varchar,
    vlr_bruto             double precision
);

alter table public.funcionarios
    owner to postgres;

