<template>
   <div 
      class="
         authFormFrame bg-fmjOrange rounded-[1.0625rem]
         plg:w-[25rem] has-[:disabled]:bg-fmjBlack/40
      "
   >
      <div
         class="
            bg-white p-10 flex-flex-col gap-4 rounded-2xl 
            shadow-[0px_8px_16px_-8px_rgba(0,_0,_0,_0.16)]
         "
      >
         <slot name="authForm"/>
      </div>

      <slot name="authAction"/>
   </div>
</template>
