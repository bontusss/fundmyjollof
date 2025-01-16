const THEME_STORAGE_KEY = "theme";

function setTheme(newTheme) {
   const theme = useCookie(THEME_STORAGE_KEY);
   theme.value = `${newTheme}`
}

async function registerTheme() {
   const theme = useCookie(THEME_STORAGE_KEY);

   //check if it's first time visit 
   if (!theme.value) {

      // set theme according to user preference
      window.matchMedia("(prefers-color-scheme: dark)").matches
         ?  setTheme('dark')
         :  setTheme('light')

   } 
}

function registerTheme(mode) {
   mode === 'dark' ? setTheme('dark') : setTheme('light')
}

export { registerTheme, registerTheme, THEME_STORAGE_KEY };
