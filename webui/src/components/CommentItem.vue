<script>
export default {
    emits: ['commentDeleted'],
    props: ["commentData", "postData"],
    data() {
        return {
            errormsg: null,
            isAuthorLoggedInUser: null
        }
    },
    mounted() {
        this.isAuthorLoggedInUser = (localStorage.getItem("userId") == this.commentData.author.userId)
    },
    methods: {
        async deleteComment() {
            this.errormsg = null
            try {
                await this.$axios.delete(`/photos/${this.postData.photoId}/comments/${this.commentData.commentId}`)
                this.$emit('comment-deleted')
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
    }
}
</script>

<template>
    <div class="card">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="card-body d-flex justify-content-between align-items-start">
            <h6 class="card-title">{{ commentData.author.username }}</h6>
            <button v-if="isAuthorLoggedInUser" type="button" class="btn btn-danger" aria-label="Close"
                @click="deleteComment">
                <svg class="feather">
                    <use href="/feather-sprite-v4.29.0.svg#trash-2" />
                </svg>
            </button>
        </div>
        <div class="card-body">
            <p class="card-text">{{ commentData.text }}</p>
            <p class="card-text">
                <small class="text-muted">{{ commentData.dateTime }}</small>
            </p>
        </div>
    </div>
</template>

<style></style>