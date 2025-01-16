import axios from '@/axios'

async function checkForUser(userInstore) {
   const indexStore = useIndexStore()

   if (!userInstore.available) {
      const { status, error, res } = await getUser()

      //if unsuccessfull
      if (status === false) {
         indexStore.pingApp({
            type:  'error',
            mssg:  'Something went wrong',
         })
      }
   }
}

async function getUser() {
   const token = useCookie('token')
   const authStore = useAuthStore()

   const result = await
      axios.get(`/user/me`, {
         headers: {
            'Authorization': `Bearer ${token.value}`
         }
      })
         .then((res) => {

            if (res.status === 200) {
               //set data in store
               authStore.user.colosach = res.data.data.user
               authStore.user.available = true

               return {
                  status: true,
                  error: null,
                  res: res.data
               }
            }
         })
         .catch((err) => {
            console.log('Error while attemting to fetch user');

            return {
               status: false,
               error: err.mssg
            }
         })

   return result
}

export {
   getUser,
   checkForUser
}