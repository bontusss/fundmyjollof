<script setup>
  definePageMeta({
      name: 'loginPage',
      layout: 'auth',
   })

   useHead({
      title: 'Fund My Jollof - Log In'
   })
   
   const buttons = reactive({
      signup:  {
         defaultText: 'Login',

         loading:  {
            state: false,
            mssg: 'Login'
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
         <template #authSecondaryAction>
            <p class="authInfo__title w-full p1 text-start"> 
               Don't have an account <br />
               <NuxtLink to="/signup" type="button" class="authInfo__link p2 inline-block text-fmjOrange" >
                  Create one
               </NuxtLink>
            </p>
            
         </template>
      </AuthCard>

      <AuthFormFrame>
         <template #authForm>
            <LoginForm />
         </template>

         <template #authAction>
            <AuthActionBtn 
               :type="`submit`"
               :text="buttons.signup.defaultText"
               :loading="buttons.signup.loading"
               :form="`loginForm`"
            />
         </template>
      </AuthFormFrame>

      <div class="flex flex-col gap-4 text-center plg:hidden">
         <div>
            <p class="authInfo__title p1"> 
               Don't have an account? 

               <NuxtLink class="authInfo__link p2  text-fmjOrange" to="/signup">
                  Create one
               </NuxtLink>
            </p>
         </div>
      </div>
   </div>
</template>
