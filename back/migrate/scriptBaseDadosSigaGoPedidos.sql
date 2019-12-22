/*

drop table rh_funcioncargos;
drop table rh_funcionarios;
drop table rh_cargos;
drop table if exists acl_usergroups;
drop table if exists acl_user;
drop table if exists acl_group;
drop table if exists aversao;

drop sequence if exists seq_acluser;
drop sequence if exists seq_aclusergroups;
drop sequence if exists seq_group;
drop sequence if exists seq_rhfuncionarios;
drop sequence if exists seq_rhcargos;
drop sequence if exists seq_rhfuncioncargos;

*/

-- drop table if exists aversao;
CREATE TABLE if not exists aversao (
    versaoid        INTEGER NOT NULL,
    versaosignota   VARCHAR(10),
    CONSTRAINT pk_aversao PRIMARY KEY ( versaoid )
);

insert into aversao
    values (1, '0.0.1') ON CONFLICT DO NOTHING;

-- drop table if exists acl_group;
CREATE TABLE IF NOT EXISTS acl_group (
    groupid          INTEGER NOT NULL,
    codigogrupo      VARCHAR(10) UNIQUE NOT NULL,
    descricaogrupo   VARCHAR(50) NOT NULL,
    CONSTRAINT pk_aclgroup PRIMARY KEY(groupid)
);

-- drop SEQUENCE if exists seq_aclgroup;
CREATE SEQUENCE if not exists seq_group;

INSERT INTO acl_group VALUES (
    nextval('seq_group'),
    'admin',
    'Administradores do sistema'
) ON CONFLICT DO NOTHING; --Garante se ao tentar inserir novamente o grupo admin, não ocorra erro e tambem não insira o usuário


-- drop table if exists acl_user
CREATE TABLE IF NOT EXISTS acl_user (
    userid          bigint NOT NULL,
    login           VARCHAR(20) unique NOT NULL,
    password        VARCHAR(50),
    datacriacao     DATE NOT NULL,
    datavalidade    DATE,
    userbloqueado   CHAR(1),
    CONSTRAINT pk_acluser PRIMARY KEY (userid)
);

-- drop sequence if exists seq_acluser;
CREATE SEQUENCE IF NOT EXISTS seq_acluser;

INSERT INTO acl_user 
    VALUES (
            nextval('seq_acluser'),
            'internal',
            MD5('Intern@l'),
            current_date,
            NULL,
            'n'
        ) ON CONFLICT DO NOTHING;

-- drop table if exists acl_usergroups;
CREATE TABLE if not EXISTS acl_usergroups (
    usergroupsid   bigint NOT NULL,
    userid         bigint NOT NULL,
    groupid        INTEGER NOT NULL,
    constraint pk_aclusergroups PRIMARY KEY (usergroupsid),
    CONSTRAINT FK_aclusergroups_aclgroup FOREIGN KEY ( groupid )
        REFERENCES acl_group ( groupid ),
    CONSTRAINT FK_aclusergroups_acluser FOREIGN KEY ( userid )
          REFERENCES acl_user ( userid )
);


create SEQUENCE if not EXISTS seq_aclusergroups;

insert into acl_usergroups
values (nextval('seq_aclusergroups'),         
        (SELECT userid            
          FROM acl_user
          WHERE login = 'internal'
         ),
         (SELECT groupid            
          FROM acl_group
          WHERE codigogrupo = 'admin'
         )) ON CONFLICT DO NOTHING;


-- 19/12/2019 
-- drop table if exists acl_usergroups;

create SEQUENCE if not EXISTS seq_rhfuncionarios;
-- drop table rh_funcionarios 
CREATE TABLE if not EXISTS rh_funcionarios (
    funcid                integer NOT NULL,
    cpf                   varchar(11)  unique,
    rg                    varchar(11),  
    funcnome              VARCHAR(40) NOT NULL,
    datanasc              DATE,
    funcdatacontratacao   DATE NOT NULL,
    funcdatadispensa      DATE    
);

ALTER TABLE IF EXISTS rh_funcionarios DROP CONSTRAINT IF EXISTS pk_rh_funcionario CASCADE;
ALTER TABLE IF EXISTS rh_funcionarios ADD CONSTRAINT pk_rh_funcionario PRIMARY KEY ( funcid );

ALTER TABLE IF EXISTS acl_user ADD IF NOT EXISTS funcid integer;
ALTER TABLE IF EXISTS acl_user
    ADD CONSTRAINT fk_acl_user_rh_funcionario FOREIGN KEY ( funcid )
        REFERENCES rh_funcionarios ( funcid );

INSERT INTO rh_funcionarios
values (nextval('seq_rhfuncionarios'), '12345678901', '54321', 'Usuário Interno com direito de Admin',
        current_date, current_date, null) ON CONFLICT DO NOTHING;

update acl_user
set funcid = (select funcid
              from rh_funcionarios
              where cpf = '12345678901')
where login = 'internal';

--22/12/2019 
create SEQUENCE if not EXISTS seq_rhcargos;
--drop table rh_cargos;
CREATE TABLE if not EXISTS rh_cargos (
    cargoid       integer NOT NULL,
    cargocodigo   VARCHAR(10) NOT NULL,
    cargonome     VARCHAR(20) NOT NULL
);

ALTER TABLE IF EXISTS rh_cargos DROP CONSTRAINT IF EXISTS  pk_rh_cargos CASCADE;
ALTER TABLE IF EXISTS rh_cargos ADD CONSTRAINT pk_rh_cargos PRIMARY KEY ( cargoid );


--drop sequence if exists seq_rhfuncioncargos;
create SEQUENCE if not EXISTS seq_rhfuncioncargos;
--drop table rh_funcioncargos;
CREATE TABLE rh_funcioncargos (
    funcionariocargoid   INTEGER NOT NULL,
    dataposse            DATE NOT NULL,
    funcid               INTEGER NOT NULL,
    cargoid              INTEGER NOT NULL
);

ALTER TABLE IF EXISTS rh_funcioncargos DROP CONSTRAINT IF EXISTS pk_rh_funcioncargos CASCADE;
ALTER TABLE IF EXISTS rh_funcioncargos ADD CONSTRAINT pk_rh_funcioncargos PRIMARY KEY ( funcionariocargoid );

ALTER TABLE IF EXISTS rh_funcioncargos
    ADD CONSTRAINT fk_rh_funcioncargos_cargos FOREIGN KEY ( cargoid )
        REFERENCES rh_cargos ( cargoid );

ALTER TABLE IF EXISTS rh_funcioncargos
    ADD CONSTRAINT fk_rh_funcioncargos_funcionarios FOREIGN KEY ( funcid )
        REFERENCES rh_funcionarios ( funcid );


commit;