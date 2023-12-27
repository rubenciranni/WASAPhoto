<script>
export default {
  data() {
    return {
      loading: false,
      errormsg: null,
      searchQuery: "",
      users: {
        records: null,
        lastId: null,
        hasNext: true
      },
    }
  },
  methods: {
    async refresh() {
      this.users = {
        records: [],
        lastId: null,
        hasNext: true
      }
      this.loadUsers()
    },
    async loadUsers() {
      this.loading = true
      try {
        let response = await this.$axios.get("/users/", { params: { username: this.searchQuery, startId: this.users.lastId } })
        if (response.data.records) {
          this.users.records = this.users.records.concat(response.data.records)
          this.users.lastId = response.data.lastId
        } else {
          this.users.hasNext = false
        }
      } catch (e) {
        if (e.response && e.response.status === 400) {
          this.errormsg = "Bad request."
        } else if (e.response && e.response.status === 401) {
          this.errormsg = "Unauthorized."
        } else if (e.response && e.response.status === 500) {
          this.errormsg = "An internal error occurred. Please try again later."
        } else {
          this.errormsg = e.toString()
        }
      } finally {
        this.loading = false
      }
    }
  },
  mounted() {
    this.$emit("logged-in")
  },
}
</script>

<template>
  <div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h2>Search</h2>
    </div>
    <div class="container mt-5">
      <form @submit.prevent="refresh">
        <div class="row">
          <div class="col-md-8">
            <input v-model="searchQuery" type="text" class="form-control" placeholder="Enter user name">
          </div>
          <div class="col-md-4">
            <button type="submit" class="btn btn-primary">Search</button>
          </div>
        </div>
      </form>
      <div class="mt-4">
        <ul class="list-group">
          <a class="list-group-item list-group-item-action" v-for="user in users.records">
            <User :user-data="user" />
          </a>
        </ul>
        <button v-if="users.records !== null && users.hasNext" @click="loadUsers" class="btn btn-primary mt-3 mb-3">Load
          More</button>
        <LoadingSpinner :loading="loading"></LoadingSpinner>
        <div v-if="!users.hasNext" class="alert alert-secondary mt-3 mb-3" role="alert">
          No more users to show.
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
      </div>
    </div>
  </div>
</template>

<style scoped></style>