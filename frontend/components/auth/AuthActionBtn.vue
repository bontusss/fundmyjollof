<script setup>
   const props = defineProps({
      text: { type: String, default: "" },
      type: { type: String, default: "" },
      path: { type: [ String, Boolean ], default: false },
      loading: { type: [ Object, Boolean ], default: false }
   })
</script>

<template>
   <button
      v-if="['button', 'submit'].includes(props.type)" 
      :type="props.type" v-bind="$attrs"
      :disabled="props.loading.state === true"
      class="
         block w-full disabled:cursor-not-allowed
         authActionBtn p-10 text-white p2 text-center 
      "
   >
      {{ props.loading.state === true ? props.loading.mssg : props.text }}
   </button>
</template>
