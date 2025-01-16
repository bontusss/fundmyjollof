import { defineStore, acceptHMRUpdate } from 'pinia'

export const useAuthStore = defineStore('auth', () => {
   const user = reactive({
      
   })

   return {
      user
   }
})

if (import.meta.hot) {
   import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot))
}
