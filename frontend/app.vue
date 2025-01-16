<script setup>
   const themeCookie = useCookie('theme')
</script>

<template>
   <div 
      class="__app" :class="{'dark': themeCookie === 'dark'}"> 
      <NuxtLayout>
         <NuxtPage />
      </NuxtLayout>
   </div>
</template>
