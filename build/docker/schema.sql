BEGIN;


CREATE TABLE IF NOT EXISTS public.account
(
    id serial NOT NULL,
    document_number character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pk_account PRIMARY KEY (id),
    CONSTRAINT account_document_number_uniquekey UNIQUE (document_number)
);

CREATE TABLE IF NOT EXISTS public.operations_type
(
    id serial NOT NULL,
    description character varying COLLATE pg_catalog."default",
    CONSTRAINT pk_operations_type PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.transaction
(
    id serial NOT NULL,
    account_id integer NOT NULL,
    operation_type_id integer NOT NULL,
    amount numeric(10, 2) NOT NULL,
    event_date timestamp without time zone NOT NULL,
    CONSTRAINT pk_transaction PRIMARY KEY (id),
	CONSTRAINT fk_operation_transaction FOREIGN KEY (operation_type_id) REFERENCES public.operations_type (id),
	CONSTRAINT fk_account_transaction FOREIGN KEY (account_id) REFERENCES public.account (id)
);

INSERT INTO public.operations_type (id,description) VALUES (1,'COMPRA A VISTA'); 
INSERT INTO public.operations_type (id,description) VALUES (2,'COMPRA PARCELADA'); 
INSERT INTO public.operations_type (id,description) VALUES (3,'SAQUE'); 
INSERT INTO public.operations_type (id,description) VALUES (4,'PAGAMENTO');

END;