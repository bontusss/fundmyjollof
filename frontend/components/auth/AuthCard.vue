<template>
   <div
      class="
         authCard flex flex-col gap-10
         plg:bg-fmjBlack/4 plg:p-10 plg:w-[17.5rem]  plg:rounded-2xl 
         plg:flex plg:flex-col plg:justify-between
      "
   >
      <h1 class="heading-1 text-fmjBlack/16 text-center plg:text-left">
         Fund my Jollof
      </h1>

      <div class="w-full flex-col gap-4 hidden plg:flex justify-end">
         <div v-if="$slots.authSecondaryAction" class="w-full flex flex-col gap-2">
            <slot name="authInfo"/>
         </div>

         <div v-if="$slots.authSecondaryAction" class="w-full flex flex-col p1">
            <slot name="authSecondaryAction"/>
         </div>
      </div>
   </div>
</template>
