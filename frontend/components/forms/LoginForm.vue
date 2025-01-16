<script setup>
   const user = reactive({
      email: {
         address: null
      },
      
      pass: {
         word: null,
         type: 'password'
      }
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
            status: false,
            type: 'info', // info or error
            mssg: '',

            default: {
               status: true,
               mssg: ''
            }
         },
         loading: false
      },

      persitUser: {
         checked: false
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
      id="loginForm"
      class="flex flex-col gap-4"
      autocomplete="off"
      @submit.prevent=""
   >
      <TextInput 
         :type="`email`"
         :group="`email-input`"
         :label="`Email address`"
         :placeholder="`Enter your email address`"
         v-model.trim="user.email.address"
         :info="inputs.email.info"
         :loading="inputs.email.loading"
         @clear-error="clearMessage(inputs.email)"
         required
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

      <div class="flex items-center justify-between gap-2 flex-wrap">
         <label class="inline-flex items-center gap-2">
            <CheckboxRoot 
               v-model:checked="inputs.persitUser.checked"
               class="
                  flex size-4 appearance-none items-center justify-center 
                  rounded-full bg-white outline-none border
               "
            >
                  <CheckboxIndicator 
                     class="
                        rounded-full text-white bg-fmjOrange size-full 
                        flex items-center justify-center p-[0.125rem]
                     "
                  >
                     <PhosCheck weight="bold"/>
                  </CheckboxIndicator>
               </CheckboxRoot>
               
               <span class="p2 text-fmjBlack/64">Remember me</span>
         </label>

         <NuxtLink
           to="/forgot-password"
           class="p2 text-fmjOrange"
         > 
            Forgot password?
         </NuxtLink>
      </div>
   </form>
</template>

<style lang="scss" scoped>

</style>