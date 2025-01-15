<script setup>
   const props = defineProps({
      group: { type: String, default: "" },
      placeholder: { type: String, default: "Input Placeholder" },
      type: { type: String, default: "text" },
      auxType: { type: [ String, Boolean ], default: false },
      label: { type: [  String, Boolean ], default: false },
      labelDescr: { type: [ String, Boolean ], default: false },
      info: { type: [ Object, Boolean ], default: false },
      loading: { type: Boolean, default: false },
      formHasError: { type: Boolean, default: false }
   })

   const model = defineModel();

   const emit = defineEmits(['clearError'])

   onMounted(() => {
    
   })
</script>

<template>
   <fieldset
      class="
         input__group pt-4 text-fmjBlack/64
         flex flex-col gap-2

         has-[:focus]:border-fmjBlack has-[:focus]:text-fmjBlack
      "
   >
      <label
         :for="props.group"
         class="caption-text"
         v-if="props.label"
      >
         {{ props.label }}
      </label>
   
      <div class="input__wrapper flex items-center justify-between gap-1 pb-4 border-b border-fmjBlack/16">
         <slot name="leftIcon"> </slot>

         <input
            :id="props.group"
            :name="props.group"
            :type="props.type"
            :placeholder="props.placeholder"
            :disabled="props.loading === true"
            v-bind="$attrs"
            v-model="model"
            autocomplete="false"
            aria-autocomplete="false"
            class="block w-full p1 bg-none autofill:bg-none font-medium"
         />
         
         <slot name="rightIcon"> </slot>
      </div>

      <span
         class="input__mssg caption-text"
         v-if="props.info.status === true"
      >
         {{ props.info.mssg }}
      </span>
   </fieldset>
</template>
