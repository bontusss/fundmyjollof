<script setup>
  definePageMeta({
      name: 'signupPage',
      layout: 'auth',
   })

   const SIGN_UP_TAB = 'signup'
   const EMAIL_VERIFICATION_TAB = 'emailVerification'

   useHead({
      title: 'Fund My Jollof - Sign Up'
   })

   const tabs = reactive({
      active: 'signup',

      signup: {

      }
   })
   
   const buttons = reactive({
      signup:  {
         defaultText: 'Create an account',

         loading:  {
            state: false,
            mssg: 'Create an account'
         },
      },
   })
  
</script>

<template>
   <div
      class="
         page-container margin w-full min-h-full 
         flex flex-col gap-10 plg:flex-row plg:w-max
      "
   >
      <AuthCard>
         <template v-if="tabs.active === 'signup'" #authInfo>
            <p class="w-full text-fmjBlack/64">
               By continuing to create an account, you accept the 
               <NuxtLink to="/terms" class="inline w-max underline text-fmjBlack"> Terms and conditions </NuxtLink> and 
               <NuxtLink to="/privacy" class="inline w-max underline text-fmjBlack"> Privacy Policy </NuxtLink> 
            </p>
         </template>

         <template #authSecondaryAction>
            <p
               class="authInfo__title w-full p1"
               v-if="tabs.active === 'signup'"
            > 
               Remember password? 
               <NuxtLink to="/login" class="authInfo__link p2 inline-block w-full text-fmjOrange" > Login </NuxtLink>
            </p>

            <p v-else class="authInfo__title w-full p1 text-start"> 
               Didn’t receive the email? <br />
               <button role="button" type="button" class="authInfo__link p2 inline-block text-fmjOrange" >
                  Resend
               </button>
            </p>
            
         </template>
      </AuthCard>

      <AuthFormFrame>
         <template #authForm>
            <SignUpForm v-if="tabs.active === 'signup'"/>   
            <EmailVerificationForm v-else/>
         </template>

         <template #authAction>
            <AuthActionBtn 
               :type="`submit`"
               :text="buttons.signup.defaultText"
               :loading="buttons.signup.loading"
               :form="`${tabs.active}Form`"
            />
         </template>
      </AuthFormFrame>

      <div class="flex flex-col gap-4 text-center plg:hidden">
         <p class="w-full text-fmjBlack/64">
            By continuing to create an account, you accept the 
            <NuxtLink to="/terms" class="inline w-max underline text-fmjBlack">Terms and conditions</NuxtLink> and 
            <NuxtLink to="/privacy" class="inline w-max underline text-fmjBlack">Privacy Policy</NuxtLink> 
         </p>

         <div>
            <p class="authInfo__title p1"> 
               Remember password? 

               <NuxtLink class="authInfo__link p2  text-fmjOrange" to="/login">
                  Login
               </NuxtLink>
            </p>
         </div>
      </div>
   </div>
</template>
