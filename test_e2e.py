from playwright.sync_api import sync_playwright

def test_search_works():
    with sync_playwright() as p:
        browser = p.chromium.launch()
        page = browser.new_page()
        
        # 1. Abrimos la app
        page.goto("http://localhost:8080")
        
        # 2. Escribimos en el buscador
        search_input = page.locator("input[name='q']")
        search_input.fill("Eslovenia")
        
        # 3. Verificamos que aparezca el resultado
        # Como usamos HTMX, esperamos a que el div se actualice
        page.wait_for_selector(".card")
        assert "Alpes" in page.content()
        
        print("✅ Prueba E2E exitosa: El buscador funciona y muestra datos.")
        browser.close()

if __name__ == "__main__":
    test_search_works()