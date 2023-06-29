-- Create table to store Brazilian states data
CREATE TABLE states (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  abbreviation CHAR(2),
  population INT
);

-- Insert data into the states table
INSERT INTO states (name, abbreviation, population) VALUES
  ('Acre', 'AC', 894470),
  ('Alagoas', 'AL', 3351543),
  ('Amapá', 'AP', 861773),
  ('Amazonas', 'AM', 4207714),
  ('Bahia', 'BA', 14873064),
  ('Ceará', 'CE', 9187103),
  ('Distrito Federal', 'DF', 3015268),
  ('Espírito Santo', 'ES', 4018650),
  ('Goiás', 'GO', 7113540),
  ('Maranhão', 'MA', 7114598),
  ('Mato Grosso', 'MT', 3526220),
  ('Mato Grosso do Sul', 'MS', 2809394),
  ('Minas Gerais', 'MG', 21168791),
  ('Pará', 'PA', 8602865),
  ('Paraíba', 'PB', 4039277),
  ('Paraná', 'PR', 11516840),
  ('Pernambuco', 'PE', 9616621),
  ('Piauí', 'PI', 3273227),
  ('Rio de Janeiro', 'RJ', 17772273),
  ('Rio Grande do Norte', 'RN', 3506853),
  ('Rio Grande do Sul', 'RS', 11422973),
  ('Rondônia', 'RO', 1855660),
  ('Roraima', 'RR', 631181),
  ('Santa Catarina', 'SC', 7252502),
  ('São Paulo', 'SP', 46649132),
  ('Sergipe', 'SE', 2398479),
  ('Tocantins', 'TO', 1607267);
