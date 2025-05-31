-- db/init/01_create_clients_table.sql
CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,                        
    cpf                  VARCHAR(18)   NOT NULL,   
    private              INTEGER       NOT NULL,   
    incompleto           INTEGER       NOT NULL,   
    data_ultima_compra   DATE,                    
    ticket_medio         NUMERIC(10,2),           
    ticket_ultima_compra NUMERIC(10,2),            
    loja_mais_frequentada TEXT,                   
    loja_ultima_compra    TEXT
)