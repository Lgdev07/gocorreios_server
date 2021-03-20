<h1 align="center">
      <img alt="GoCorreiosServer" title="GoCorreiosServer" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  GoCorreios Server
</h3>

<p align="center">Api to serve correios services in a easy way 📖</p>
<p align="center">Made with Golang 🚀</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/gocorreios_server?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/gocorreios_server/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/gocorreiosweb?style=social">
  </a>
</p>

<p align="center">
  <a href="#-installation-and-execution">Installation and execution</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-available-routes">Available Routes</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;
</p>


## 🚀 Installation and execution

1. Clone this repository and go to the directory;
2. Rename sample .env;

<h4> 🔧 Development </h4>

1. Run `docker-compose up`;
2. Make the Requests;

<h4> 🧪 Tests </h4>

1. Run `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`;

## 🛣️ Available Routes

  POST - Create Fare:
  - **/fares** <br>
  Expected Json Body Request:<br>
  ```
  {
    "service": "SEDEX",
		"cep_destination": "09981380",
		"cep_origin": "09981380",
		"height": "15",
		"lenght": "15",
		"weight": "1",
		"width": "15",
  }
  ```
  Expected Json Response:<br>
  ```
  {
    "service": "SEDEX",
		"price": "22,50",
		"days_for_delivery": "3",
		"deliver_home": "N",
		"deliver_saturday": "S",
		"obs": "O CEP de destino está temporariamente sem entrega domiciliar. A entrega será efetuada na agência indicada no Aviso de Chegada que será entregue no endereço do destinatário."
  }
  ```

  POST - Create Tracking:
  - **/trackings** <br>
  Expected Json Body Request:<br>
  ```
  {
    "codes": ["ON732404576BR", "ON732404576BR"]
  }
  ```
  Expected Json Response:<br>
  ```
[
  {
    "category": "SEDEX",
    "error": "",
    "history": [
      {
        "date": "03/03/2021 - 14:38",
        "description": "Objeto em trânsito - por favor aguarde",
        "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
        "detail": "",
        "origin": "AGF SHOPPING CAMPO LIMPO - ESTRADA DO CAMPO LIMPO, 459, PIRAJUSSARA - SAO PAULO/SP",
        "status": 1,
        "type": "RO"
      },
      {
        "date": "03/03/2021 - 14:05",
        "description": "Objeto postado",
        "destination": "",
        "detail": "",
        "origin": "AGF SHOPPING CAMPO LIMPO - ESTRADA DO CAMPO LIMPO, 459, PIRAJUSSARA - SAO PAULO/SP",
        "status": 1,
        "type": "PO"
      }
    ],
    "last_date": "05/03/2021 - 14:04",
    "last_description": "Objeto entregue ao destinatário",
    "last_destination": "",
    "last_detail": "",
    "last_origin": "CDD EXEMPLO",
    "last_status": 1,
    "last_type": "BDE",
    "number": "ON732404576BR"
  },
  {
    "category": "SEDEX",
    "error": "",
    "history": [
      {
        "date": "03/03/2021 - 14:38",
        "description": "Objeto em trânsito - por favor aguarde",
        "destination": "CTE SAUDE - RUA DO BOQUEIRAO, 320, SAUDE - SAO PAULO/SP",
        "detail": "",
        "origin": "AGF SHOPPING CAMPO LIMPO - ESTRADA DO CAMPO LIMPO, 459, PIRAJUSSARA - SAO PAULO/SP",
        "status": 1,
        "type": "RO"
      },
      {
        "date": "03/03/2021 - 14:05",
        "description": "Objeto postado",
        "destination": "",
        "detail": "",
        "origin": "AGF SHOPPING CAMPO LIMPO - ESTRADA DO CAMPO LIMPO, 459, PIRAJUSSARA - SAO PAULO/SP",
        "status": 1,
        "type": "PO"
      }
    ],
    "last_date": "05/03/2021 - 14:04",
    "last_description": "Objeto entregue ao destinatário",
    "last_destination": "",
    "last_detail": "",
    "last_origin": "CDD EXEMPLO",
    "last_status": 1,
    "last_type": "BDE",
    "number": "ON732404576BR"
  }
]
  ```

## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---
