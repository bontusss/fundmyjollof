<script setup>
   const indexStore = useIndexStore()
   const DOMStore = useDOMStore()
   const route = useRoute()
</script>

<template>
   <div class="layout-container">
      <slot />
   </div>
</template>