<template>
  <div>
    {{ data.data }}
  </div>
</template>

<script setup>
const data = await useFetch('http://localhost:8060');

</script>
