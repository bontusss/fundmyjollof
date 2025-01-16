<script setup>
   const user = reactive({
      email: {
         address: null,
         sent: false
      },
      
      pass: {
         word: null,
         type: 'password'
      },
      
      onboard: true
   })

   const inputs = reactive({
      email: {
         info: {
            status: false,
            type: 'info', // info or error
            mssg: '',

            default: {
               status: false,
               mssg: ''
            }
         },
         loading: false
      },

      password: {
         info: {
            status: true,
            type: 'info', // info or error
            mssg: 'Your password must be at least 8 characters long and include a combination of uppercase and lowercase letters, at least one numeric digit, and at least one special character ',

            default: {
               status: true,
               mssg: 'Your password must be at least 8 characters long and include a combination of uppercase and lowercase letters, at least one numeric digit, and at least one special character '
            }
         },
         loading: false
      }
   }) 

   function togglePassVisibility() {
      user.pass.type === 'password' 
         ? user.pass.type = 'text' 
         : user.pass.type = 'password'
   }
</script>

<template>
   <form
      id="signUpForm"
      class="flex flex-col gap-4"
      autocomplete="off"
      @submit.prevent=""
   >
      <TextInput 
         :type="`email`"
         :group="`email-input`"
         :label="`Email address`"
         :placeholder="`Enter your email address`"
         v-model="user.email.address"
      />

      <TextInput 
         :type="user.pass.type" 
         :group="`password-input`"
         :label="`Password`"
         :placeholder="`Enter your password`"
         v-model.trim="user.pass.word"
         :info="inputs.password.info"
         :loading="inputs.password.loading"
         @clear-error="clearMessage(inputs.password)"
         required
      >
         <template #rightIcon>
            <button 
               type="button" role="button" 
               @click="togglePassVisibility"
            >
               <PhosEyeSlash weight="bold" v-if="user.pass.type === 'password'"/>
               <PhosEye weight="bold" v-else/>
            </button> 
         </template>  
      </TextInput>


   </form>
</template>