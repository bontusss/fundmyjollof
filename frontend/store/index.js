import { defineStore, acceptHMRUpdate } from 'pinia'

export const useIndexStore = defineStore('index', () => {

   const app = reactive({
      hydrated: false
   })
   const toast = reactive({
      isShown: false,
      type: 'info', // error, success, info
      mssg: 'Toast mssg'
   })


   const pingApp = (payload) => {
      toast.type = payload.type
      toast.mssg = payload.mssg
      toast.isShown = true 

      //hide after 3 secs
      setTimeout(() => {
         toast.isShown = false    
      }, 3000);

      //reset after 5 secs
      setTimeout(() => {
         toast.type = 'info'
         toast.mssg = 'Toast mssg'
      }, 5000);
   }

   const openModal = (name) => {
      modal[name].isShown = true
      modal.isActive = true
      modal.name = name
   }

   const closeModal = (name) => {
      modal[modal.name].isShown = false
      modal.isActive = false
      modal.name = null
   }

   const toggleMenu = () => {
      menu.isShown = !menu.isShown
   }

   return {
      app, menu, preloader, modal, toast, featuredArtists, pingApp,
      openModal, closeModal, toggleMenu 
   }
})

if (import.meta.hot) {
   import.meta.hot.accept(acceptHMRUpdate(useIndexStore, import.meta.hot))
}
