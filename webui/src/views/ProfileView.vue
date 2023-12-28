<script>
export default {
    data: function () {
        return {
            errormsg: {
                userNotFound: false,
                profile: null,
                photos: null
            },
            loading: {
                profile: false,
                photos: false,
            },
            isUserLoggedInUser: null,
            newUsername: null,
            user: {
                username: null,
                userId: null
            },
            photos: {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true
            },
            followers: {
                records: [],
                lastId: null,
                hasNext: true
            },
            following: {
                records: [],
                lastId: null,
                hasNext: true
            }
        }
    },
    created() {
        this.$watch(
            () => this.$route.params,
            (toParams, previousParams) => {
                location.reload()
            }
        )
    },
    async mounted() {
        this.$emit("logged-in")
        await this.getUser()
        this.isUserLoggedInUser = (localStorage.getItem("userId") == this.user.userId)
        if (!this.user.username) {
            this.errormsg.userNotFound = true
        } else {
            await this.loadProfile()
            this.loadPhotos()
        }
    },
    methods: {
        async getUser() {
            this.errormsg.profile = null
            try {
                let response = await this.$axios.get("/users/", { params: { username: this.$route.params.username } })
                if (response.data.records.length == 1) {
                    this.user = response.data.records[0]
                } else {
                    this.errormsg.profile = "Not Found."
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            }
        },
        async loadProfile() {
            this.loading.profile = true
            this.errormsg.profile = null
            try {
                let response = await this.$axios.get(`/users/${this.user.userId}`)
                this.user = response.data
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg.profile = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            } finally {
                this.loading.profile = false
            }
        },
        async loadPhotos() {
            this.loading.photos = true
            this.errormsg.photos = null
            try {
                let response = await this.$axios.get("/photos/", {
                    params: {
                        startDate: this.photos.lastDate,
                        startId: this.photos.lastId,
                        userId: this.user.userId
                    }
                })
                if (response.data.records) {
                    this.photos.records = this.photos.records.concat(response.data.records)
                    this.photos.lastDate = response.data.lastDate
                    this.photos.lastId = response.data.lastId
                    if (response.data.records.length < this.$paginationLimit) {
                        this.photos.hasNext = false
                    }
                } else {
                    this.photos.hasNext = false
                }

            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.photos = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.photos = "Unauthorized."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.photos = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.photos = e.toString()
                }
            } finally {
                this.loading.photos = false
            }
        },
        async loadFollowers() {
            this.loading.profile = true
            this.errormsg.profile = null
            try {
                let response = await this.$axios.get(`/users/${this.user.userId}/followers/`, {
                    params: {
                        startId: this.followers.lastId
                    }
                })
                if (response.data.records) {
                    this.followers.records = this.followers.records.concat(response.data.records)
                    this.followers.lastId = response.data.lastId
                    if (response.data.records.length < this.$paginationLimit) {
                        this.followers.hasNext = false
                    }
                } else {
                    this.followers.hasNext = false
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg.profile = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            } finally {
                this.loading.profile = false
            }
        },
        resetFollowers() {
            this.followers = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            }
        },
        async loadFollowing() {
            this.loading.profile = true
            this.errormsg.profile = null
            try {
                let response = await this.$axios.get(`/users/${this.user.userId}/following/`, {
                    params: {
                        startId: this.following.lastId
                    }
                })
                if (response.data.records) {
                    this.following.records = this.following.records.concat(response.data.records)
                    this.following.lastId = response.data.lastId
                    if (response.data.records.length < this.$paginationLimit) {
                        this.following.hasNext = false
                    }
                } else {
                    this.following.hasNext = false
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg.profile = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            } finally {
                this.loading.profile = false
            }
        },
        resetFollowing() {
            this.following = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            }
        },
        async toggleFollow() {
            this.errormsg.profile = null
            try {
                if (!this.user.isFollowed) {
                    await this.$axios.put(`/following/${this.user.userId}`)
                    this.user.numberOfFollowers++
                    this.user.isFollowed = true
                } else {
                    await this.$axios.delete(`/following/${this.user.userId}`)
                    this.user.numberOfFollowers--
                    this.user.isFollowed = false
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg.profile = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            }
        },
        async toggleBan() {
            this.errormsg.profile = null
            try {
                if (!this.user.isBanned) {
                    await this.$axios.put(`/bans/${this.user.userId}`)
                    this.user.isBanned = true
                } else {
                    await this.$axios.delete(`/bans/${this.user.userId}`)
                    this.user.isBanned = false
                }
            } catch (error) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg.profile = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            }
        },
        async updateUsername() {
            this.errormsg.profile = null
            try {
                await this.$axios.put("settings/username", { newUsername: this.newUsername })
                this.user.username = this.newUsername
                localStorage.setItem("username", this.newUsername)
                this.$router.replace(`/${this.newUsername}`)
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg.profile = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg.profile = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg.profile = "New username is already taken."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg.profile = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg.profile = e.toString()
                }
            }
        },
        handlePostDeleted() {
            this.photos = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true
            }
            this.loadPhotos()
        },
        isUsernameValid() {
            const usernameRegex = /^[a-zA-Z0-9_-]{3,16}$/
            return usernameRegex.test(this.newUsername)
        }
    }
}
</script>

<template>
    <div v-if="!errormsg.userNotFound">
        <!-- Profile info -->
        <div class="container pt-3 pb-2 mt-3 mb-3 border-bottom">
            <h2 class=""> {{ user.username }}</h2>
            <ErrorMsg v-if="errormsg.profile" :msg="errormsg.profile"></ErrorMsg>
            <div class="row mb-3">
                <div class="col-2">
                    <span class="ml-2">
                        {{ user.numberOfPhotos }}
                        <strong>Posts</strong>
                    </span>
                </div>
                <div class="col-2">
                    <UserListModal title="Followers" :users-data="followers" @reset-users="resetFollowers"
                        @update-users="loadFollowers"></UserListModal>
                    <span>
                        {{ user.numberOfFollowers }}
                        <strong>Followers</strong>
                    </span>
                </div>
                <div class="col-2">
                    <UserListModal title="Following" :users-data="following" @reset-users="resetFollowing"
                        @update-users="loadFollowing"></UserListModal>
                    <span>
                        {{ user.numberOfFollowing }}
                        <strong>Following</strong>
                    </span>
                </div>
            </div>
            <div v-if="!isUserLoggedInUser" class="flex-row mb-2">
                <button @click="toggleFollow" type="button" class="btn btn-primary me-2">
                    <svg v-if="!user.isFollowed" class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#user-plus" />
                    </svg>
                    <svg v-else class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#user-minus" />
                    </svg>
                    {{ user.isFollowed ? "Unfollow" : "Follow" }}
                </button>
                <button @click="toggleBan" type="button"
                    :class="!user.isBanned ? 'btn btn-danger me-2' : 'btn btn-success'">
                    <svg v-if="!user.isBanned" class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#user-x" />
                    </svg>
                    <svg v-else class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#user-check" />
                    </svg>
                    {{ user.isBanned ? "Unban" : "Ban" }}
                </button>
            </div>
            <div v-else>
                <button type="button" data-bs-toggle="modal" data-bs-target="#setUsernameModal" class="btn btn-secondary">
                    <svg class="feather">
                        <use href="/feather-sprite-v4.29.0.svg#edit" />
                    </svg>
                    Set new username
                </button>
                <div class="modal fade" id="setUsernameModal" tabindex="-1" aria-labelledby="setUsernameModalLabel"
                    aria-hidden="true">
                    <div class="modal-dialog modal-dialog-centered modal-lg">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h5 class="modal-title" id="setUsernameModalLabel">Set new username</h5>
                                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                            </div>
                            <form @submit.prevent="updateUsername">
                                <div class="modal-body">
                                    <div class="mb-3">
                                        <label for="username" class="form-label">Username</label>
                                        <input type="text" class="form-control" id="username"
                                            aria-describedby="usernameHelp" v-model="newUsername"
                                            :class="{ 'is-invalid': !isUsernameValid() }">
                                        <div id="usernameHelp" class="form-text">
                                            Your username must be 3-16 characters long, containing only letters (a-z, A-Z),
                                            numbers (0-9),
                                            hyphens (-), and underscores (_).
                                        </div>
                                    </div>
                                </div>
                                <div class="modal-footer">
                                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                                    <button :disabled="!newUsername || !isUsernameValid()" type="submit"
                                        class="btn btn-primary" data-bs-dismiss="modal">Submit</button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
            <LoadingSpinner :loading="loading.profile"></LoadingSpinner>
        </div>
        <!-- Posts -->
        <ErrorMsg v-if="errormsg.photos" :msg="errormsg.photos"></ErrorMsg>
        <PostList :posts-data="photos" @load-more="loadPhotos" @post-deleted="handlePostDeleted"></PostList>
        <LoadingSpinner :loading="loading.photos"></LoadingSpinner>
    </div>
    <div v-else>
        <ErrorMsg class="mt-3" msg="User not found."></ErrorMsg>
    </div>
</template>

<style scoped></style>