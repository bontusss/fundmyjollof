import { defineStore, acceptHMRUpdate } from 'pinia'

export const useDOMStore = defineStore('dom', () => {

   return {
   }
})

if (import.meta.hot) {
   import.meta.hot.accept(acceptHMRUpdate(useDOMStore, import.meta.hot))
}
