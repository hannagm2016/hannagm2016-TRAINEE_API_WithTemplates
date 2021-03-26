<template>
  <div id="app">
    <router-link to="post/66">Post</router-link>
    <router-link to="edit/66">Edit</router-link>
    <router-link to="create">Create</router-link>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
</style>
