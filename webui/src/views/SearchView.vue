<script>
export default {
  data() {
    return {
      searchQuery: "",
      users: [],
      pagination: {
        lastId: null,
      },
    }
  },
}
</script>

<template>
    <div>
        <div
			  class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
    		<h2>Search</h2>
  		</div>
    </div>
</template>

<style scoped>
</style>