/*
drop table if exists acl_usergroups;
drop table if exists acl_user;
drop table if exists acl_group;
drop table if exists aversao;
drop sequence if exists seq_acluser;
drop sequence if exists seq_aclusergroups
drop sequence if exists seq_group

*/

-- drop table if exists aversao;
CREATE TABLE if not exists aversao (
    versaoid        INTEGER NOT NULL,
    versaosignota   VARCHAR(10),
    CONSTRAINT pk_aversao PRIMARY KEY ( versaoid )
);

insert into aversao
    values (1, '0.0.1') ON CONFLICT DO NOTHING;;

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
         ));




commit;




