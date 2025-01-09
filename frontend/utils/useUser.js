import axios from '@/axios'

async function CHECK_FOR_USER(userInstore) {
   const indexStore = useIndexStore()

   if (!userInstore.available) {
      const { status, error, res } = await GET_USER()

      //if unsuccessfull
      if (status === false) {
         indexStore.PING_APP({
            type:  'error',
            mssg:  'Something went wrong',
         })
      }
   }
}

async function GET_USER() {
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
   GET_USER,
   CHECK_FOR_USER
}