<template>
  <div>
    {{ res.data }}
  </div>
</template>

<script setup>
const res = await useFetch('http://192.168.1.35/api')
console.log(res.data);
</script>