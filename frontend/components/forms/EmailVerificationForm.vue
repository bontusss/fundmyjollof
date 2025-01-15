<script setup>
   const OTP = ref([])
   const emailVerificationForm = ref(null)

   function verifyEmail() {
      console.log('Email verification form submitted')
   }
</script>

<template>
   <form
      autocomplete="off"
      @submit.prevent="verifyEmail"
      id="emailVerificationForm"
      ref="emailVerificationForm"
      class="flex flex-col gap-4"
   >
      <h2 class="title-text text-fmjBlack"> Verify your email address</h2>

      <p class="p1 text-fmjBlack/64">
         A verification code has been sent to your email. 
         Please enter the code below to complete your registration.
         <br />
         <br />
         The code will expire in 10 minutes.
      </p>

      <div>
         <PinInputRoot
            id="pin-input"
            v-model="OTP"
            placeholder="-"
            class="flex gap-4 items-center justify-between"
            @complete="verifyEmail"
         >
            <PinInputInput
               v-for="(id, index) in 5"
               :key="id"
               :index="index"
               class="
                  p1 py-4 w-full bg-white text-center border-b border-fmjBlack/16
                  focus:border-fmjOrange text-fmjOrange
               "
               required
            />
         </PinInputRoot>
      </div>
   </form>
</template>
