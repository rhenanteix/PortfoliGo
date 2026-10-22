# PortifoliGo

**Motor open source de processamento e análise de carteiras de recebíveis desenvolvido em Go (Golang).**

O **PortifoliGo** é um projeto educacional de backend, desenvolvido com código aberto, que tem como objetivo explorar a linguagem Go, o processamento de dados financeiros, a validação da qualidade dos dados, a arquitetura modular e a análise baseada em regras de negócio.

## Visão Geral

O PortifoliGo processa carteiras de recebíveis fornecidas por meio de arquivos CSV, padroniza suas estruturas, identifica possíveis inconsistências e gera indicadores financeiros para facilitar a análise das informações.

O projeto prioriza o **processamento determinístico de dados**, utilizando regras explícitas e verificáveis, em vez de decisões automatizadas de crédito.

Seu principal objetivo é demonstrar, na prática, conceitos de engenharia de software aplicados a um cenário financeiro realista.

## Funcionalidades Principais

- Leitura e importação de arquivos CSV.
- Mapeamento automático de colunas utilizando nomes e aliases conhecidos.
- Normalização de valores monetários, documentos e datas.
- Validação de CPF e CNPJ.
- Identificação de possíveis registros duplicados.
- Tratamento de erros individualmente por linha, sem interromper todo o processamento.
- Geração de resumos sobre a qualidade dos dados.
- Consolidação e análise de carteiras de recebíveis.
- Cálculo de indicadores como valor total, quantidade de títulos, quantidade de sacados e ticket médio.
- Análise de concentração por sacado.
- Cálculo de prazo médio ponderado e distribuição de vencimentos.
- Geração de insights e alertas baseados em regras configuráveis.

**Observação:** algumas funcionalidades avançadas de mapeamento e geração de insights ainda estão em desenvolvimento e aprimoramento.

## Arquitetura

O backend utiliza uma arquitetura modular, organizada em pacotes independentes, com responsabilidades bem definidas.

Essa abordagem facilita a manutenção, a realização de testes automatizados e a evolução do projeto.

### Organização dos módulos

- `parser` — Leitura e interpretação dos arquivos de entrada.
- `mapping` — Identificação e correspondência das colunas com os campos internos do sistema.
- `normalization` — Padronização e conversão dos dados recebidos.
- `validation` — Aplicação de regras para verificar a qualidade e a consistência dos dados.
- `processing` — Orquestração do processamento e tratamento individual dos registros.
- `analysis` — Cálculo de indicadores e métricas das carteiras.
- `insights` — Geração de análises e alertas a partir de regras configuráveis.
- `scan` — Gerenciamento dos endpoints HTTP relacionados ao processamento.

### Fluxo de processamento

1. **Importação:** o usuário envia um arquivo CSV contendo uma carteira de recebíveis.
2. **Leitura:** o sistema interpreta o arquivo e identifica suas colunas e registros.
3. **Mapeamento:** as colunas são associadas aos campos internos do PortifoliGo.
4. **Normalização:** os valores, documentos e datas são convertidos para formatos padronizados.
5. **Validação:** o sistema identifica inconsistências, documentos inválidos e possíveis duplicidades.
6. **Processamento:** os registros são tratados individualmente, permitindo que problemas pontuais não interrompam toda a análise.
7. **Análise:** são calculados indicadores financeiros e estatísticos da carteira.
8. **Insights:** regras de negócio identificam condições que merecem atenção.
9. **Resultado:** a API retorna os dados processados, indicadores e apontamentos em formato JSON.

## Tecnologias Utilizadas

- **Go (Golang)** — Linguagem principal do backend.
- **Biblioteca padrão do Go** — Utilização de pacotes como `net/http`, `encoding/csv` e `encoding/json`.
- **Google UUID** — Geração de identificadores únicos.
- **Go Testing** — Desenvolvimento e execução de testes automatizados.

O projeto busca aproveitar os recursos nativos do Go e manter a quantidade de dependências externas reduzida.

## Como Executar o Projeto

### Pré-requisitos

É necessário ter o **Go instalado** no ambiente de desenvolvimento.

### Instalação e execução

Clone o repositório e acesse a pasta do projeto.

Em seguida, execute:

```bash
go mod download
go run ./cmd/api
```

A API estará disponível em:

`http://localhost:8080`

### Importação de um arquivo CSV

Para testar o processamento de uma carteira, utilize o seguinte comando:

```bash
curl -F "file=@examples/carteira.csv" \
  http://localhost:8080/api/v1/scans/upload
```

O sistema receberá o arquivo, executará as etapas de processamento e retornará os resultados em JSON.

### Executar os testes automatizados

```bash
go test ./...
```

### Verificar o código

```bash
go vet ./...
```

### Formatar o código

```bash
gofmt -w .
```

## Roadmap — Próximas Evoluções

O PortifoliGo está em desenvolvimento contínuo. Entre as melhorias planejadas estão:

- [ ] Concluir o fluxo de mapeamento manual de colunas.
- [ ] Implementar modelos reutilizáveis de mapeamento para diferentes layouts.
- [ ] Adicionar suporte a arquivos Excel (`.xlsx`).
- [ ] Implementar persistência de dados com PostgreSQL.
- [ ] Criar histórico de análises e arquivos processados.
- [ ] Implementar processamento assíncrono para grandes volumes de dados.
- [ ] Gerar relatórios em PDF e Excel.
- [ ] Desenvolver um dashboard web para visualização dos indicadores.
- [ ] Criar uma central de revisão e correção de inconsistências.
- [ ] Permitir a comparação entre diferentes carteiras.
- [ ] Implementar autenticação, controle de acesso e registros de auditoria.
- [ ] Ampliar a cobertura de testes automatizados.

## Objetivos de Aprendizado

O PortifoliGo foi criado como um ambiente prático de estudos e experimentação em engenharia de software.

Os principais conceitos explorados são:

- Desenvolvimento idiomático com Go.
- Organização de projetos utilizando módulos e pacotes.
- Modelagem de domínio e estruturas de dados.
- Composição e utilização de interfaces.
- Tratamento explícito de erros.
- Desenvolvimento de APIs HTTP.
- Leitura e processamento de arquivos.
- Normalização e validação de dados.
- Implementação de regras de negócio determinísticas.
- Desenvolvimento de testes unitários.
- Separação de responsabilidades e arquitetura modular.
- Construção de aplicações backend extensíveis e de fácil manutenção.

O objetivo é evoluir o projeto gradualmente, aplicando boas práticas de desenvolvimento e explorando novos recursos da linguagem Go.

## Aviso Importante

O **PortifoliGo é um projeto educacional e experimental**.

A ferramenta não realiza aprovação ou reprovação de crédito, não fornece recomendações de investimento e não substitui a análise de profissionais especializados.

Os indicadores e alertas gerados têm caráter informativo e destinam-se ao estudo de processamento e análise de dados.

## Licença

O projeto poderá ser disponibilizado sob uma licença open source permissiva, como **MIT** ou **Apache 2.0**, antes de sua publicação oficial.

A licença escolhida deverá ser incluída em um arquivo `LICENSE` na raiz do repositório.