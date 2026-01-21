# 🎌 Países Confusos - Estilo 1973

[![CI-CD_Paises_Confusos]https://github.com/santiagourdaneta/Web-App-Paises-Confusos-HTMLFirst-ZeroNode-GoLang-MinimalistUI/actions/workflows/ci.yml/badge.svg)](https://github.com/santiagourdaneta/Web-App-Paises-Confusos-HTMLFirst-ZeroNode-GoLang-MinimalistUI/actions)

Una aplicación web **User-First** diseñada para laptops y celulares antiguos, donde la velocidad y la estética se encuentran. Inspirada en la elegancia letal de **Meiko Kaji** en *Lady Snowblood (1973)*.

## 🧠 Filosofía del Proyecto

Este proyecto rechaza el "bloatware" de la web moderna. Mientras que una página promedio pesa 2MB, esta aplicación entrega interactividad instantánea con menos de 20KB de transferencia inicial.

- **Zero-Node:** Construido íntegramente en Go para máxima eficiencia de CPU.
- **HTML-First:** Interactividad mediante **HTMX**, eliminando la necesidad de pesados archivos JS en el cliente.
- **1973 Aesthetic:** UI inspirada en el cine Noir japonés: rojos sangre, negros profundos y una tipografía que evoca los créditos de una película de culto.

## 🛠️ Stack Tecnológico

| Componente | Tecnología | Razón |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) | Velocidad de ejecución y binario único. |
| **Frontend** | HTMX / CSS3 | Interactividad sin JS complejo. |
| **Infraestructura** | Docker (Alpine) | Imagen final de solo 15MB. |
| **Calidad** | TDD | Pruebas unitarias, integracion, estrés y E2E sin dependencias de Node. |

## 🔐 Seguridad y SEO

- **Encabezados de Seguridad:** Implementación de CSP, XSS Protection y Anti-Clickjacking.
- **Rate Limiting:** Protección nativa contra abusos y DDOS básico.
- **SEO Ready:** Meta-tags OpenGraph configurados para visibilidad en LinkedIn, Twitter y WhatsApp.
- **Accesibilidad:** Contraste de color validado y marcado semántico HTML5.

## 🎵 Inspiración Cultural

El diseño sonoro y visual rinde homenaje a la era **Showa** de Japón. La estructura minimalista del código refleja la actuación de **Meiko Kaji**, donde "menos es más". 

## 🚀 Instalación y Uso

```bash
# Clonar y entrar
git clone [https://github.com/santiagourdaneta/Web-App-Paises-Confusos-HTMLFirst-ZeroNode-GoLang-MinimalistUI](https://github.com/santiagourdaneta/Web-App-Paises-Confusos-HTMLFirst-ZeroNode-GoLang-MinimalistUI)
cd Web-App-Paises-Confusos-HTMLFirst-ZeroNode-GoLang-MinimalistUI

# Correr con Makefile
make run

