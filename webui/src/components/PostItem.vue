<script>
export default {
    emits: ['postDeleted', 'commentAdded', 'commentDeleted', 'likeAdded', 'likeDeleted'],
    props: ["postData"],
    data() {
        return {
            errormsg: null,
            loading: false,
            isAuthorLoggedInUser: null,
            comments: {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            },
            likes: {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            },
            newCommentText: ""
        }
    },
    mounted() {
        this.isAuthorLoggedInUser = (localStorage.getItem("userId") == this.postData.author.userId)
    },
    methods: {
        async addComment() {
            this.loading = true
            this.errormsg = null
            try {
                await this.$axios.post(`/photos/${this.postData.photoId}/comments/`, { text: this.newCommentText })
                this.reloadComments()
                this.$emit('comment-added')
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg = e.toString()
                }
            }
            this.loading = false
        },
        async loadComments() {
            this.loading = true
            this.errormsg = null
            try {
                let response = await this.$axios.get(`/photos/${this.postData.photoId}/comments/`, {
                    params: {
                        startDate: this.comments.lastDate,
                        startId: this.comments.lastId,
                    }
                })
                if (response.data.records) {
                    this.comments.records = this.comments.records.concat(response.data.records)
                    this.comments.lastDate = response.data.lastDate
                    this.comments.lastId = response.data.lastId
                    if (response.data.records.length < this.$paginationLimit) {
                        this.comments.hasNext = false
                    }
                } else {
                    this.comments.hasNext = false
                }

            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "Forbidden."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg = e.toString()
                }
            }
            this.loading = false
        },
        reloadComments() {
            this.comments = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            }
            this.newCommentText = ""
            this.loadComments()
        },
        resetComments() {
            this.comments = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            }
        },
        async handleCommentDeleted() {
            this.reloadComments()
            this.$emit('comment-deleted')
        },
        async toggleLike() {
            try {
                if (!this.postData.isLiked) {
                    await this.$axios.put(`/liked-photos/${this.postData.photoId}`)
                    this.resetLikes()
                    this.$emit('like-added')
                } else {
                    await this.$axios.delete(`/liked-photos/${this.postData.photoId}`)
                    this.resetLikes()
                    this.$emit('like-deleted')
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg = "Not found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg = e.toString()
                }
            }
        },
        async loadLikes() {
            this.loading = true
            this.errormsg = null
            try {
                let response = await this.$axios.get(`/photos/${this.postData.photoId}/likes/`, {
                    params: {
                        startId: this.likes.lastId,
                    }
                })
                if (response.data.records) {
                    this.likes.records = this.likes.records.concat(response.data.records)
                    this.likes.lastDate = response.data.lastDate
                    this.likes.lastId = response.data.lastId
                    if (response.data.records.length < this.$paginationLimit) {
                        this.likes.hasNext = false
                    }
                } else {
                    this.likes.hasNext = false
                }

            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "Forbidden."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg = e.toString()
                }
            } finally {
                this.loading = false
            }
        },
        resetLikes() {
            this.likes = {
                records: [],
                lastDate: null,
                lastId: null,
                hasNext: true,
            }
        },
        async deletePost() {
            try {
                await this.$axios.delete(`/photos/${this.postData.photoId}`)
                this.$emit('post-deleted')
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Bad request."
                } else if (e.response && e.response.status === 401) {
                    this.errormsg = "Unauthorized."
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "Forbidden."
                } else if (e.response && e.response.status === 404) {
                    this.errormsg = "Not Found."
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. Please try again later."
                } else {
                    this.errormsg = e.toString()
                }
            }
        }
    },
}
</script>

<template>
    <div class="container mt-3 mb-3">
        <div class="row">
            <div class="col-md-5 mx-auto">
                <div class="card">
                    <img :src="`${$axios.defaults.baseURL}/photos/${postData.photoId}`" class="card-img-top"
                        alt="Post Image" />
                    <div class="card-body">
                        <a :href="`#/${postData.author.username}`">
                            <h6 class="card-title">{{ postData.author.username }}</h6>
                        </a>
                        <p class="card-text">{{ postData.caption }}</p>
                        <p class="card-text">
                            <small class="text-muted">{{ postData.dateTime }}</small>
                        </p>
                    </div>
                    <div class="card-footer">
                        <div class="flex-row">
                            <UserListModal title="Likes" :users-data="likes" @reset-users="resetLikes"
                                @update-users="loadLikes"></UserListModal>
                            <button @click="toggleLike" type="button" class="btn btn-primary me-2">
                                <svg class="feather">
                                    <use href="/feather-sprite-v4.29.0.svg#heart" />
                                </svg>
                                {{ postData.isLiked ? "Unlike" : "Like" }}
                                <span class="badge badge-light">
                                    {{ postData.numberOfLikes }}
                                </span>
                            </button>
                            <button @click="loadComments" type="button" data-bs-toggle="modal"
                                :data-bs-target="'#commentModal' + postData.photoId" class="btn btn-secondary me-2">
                                <svg class="feather">
                                    <use href="/feather-sprite-v4.29.0.svg#message-circle" />
                                </svg>
                                Comment
                                <span class="badge badge-light"> {{ postData.numberOfComments }}</span>
                            </button>
                            <button v-if="isAuthorLoggedInUser" type="button" class="btn btn-danger" aria-label="Close"
                                @click="deletePost">
                                <svg class="feather">
                                    <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                                </svg>
                            </button>
                        </div>
                        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="modal fade" :id="'commentModal' + postData.photoId" tabindex="-1"
        aria-labelledby="'commentModalLabel' + postData.photoId" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-lg">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" :id="'commentModalLabel' + postData.photoId">Comments</h5>
                    <button @click="resetComments" type="button" class="btn-close" data-bs-dismiss="modal"
                        aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <ul class="list-group">
                        <li class="list-group-item" v-for="comment in comments.records" :key="comment.commentId">
                            <CommentItem @comment-deleted="handleCommentDeleted" :comment-data="comment"
                                :post-data="postData" />
                        </li>
                    </ul>
                    <div class="text-center">
                        <button v-if="comments.hasNext" @click="loadComments" class="btn btn-primary mt-3 mb-3">Load
                            More</button>
                        <LoadingSpinner :loading="loading"></LoadingSpinner>
                        <div v-if="!comments.hasNext" class="alert alert-secondary mt-3 mb-3" role="alert">
                            No more comments to show.
                        </div>
                    </div>
                    <div class="mt-3">
                        <label for="commentInput" class="form-label">Add a Comment:</label>
                        <textarea v-model="newCommentText" class="form-control" id="commentInput" rows="3"></textarea>
                        <button @click="addComment" type="button" class="btn btn-primary mt-2">Add Comment</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style></style>
