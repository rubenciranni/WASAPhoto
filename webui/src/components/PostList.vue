<script>
export default {
    props: ['postsData'],
    emits: ['postDeleted', 'loadMore'],
    methods: {
        handlePostDeleted() {
            this.$emit('post-deleted')
        },
        handleCommentAdded(post) {
            post.numberOfComments++
        },
        handleCommentDeleted(post) {
            post.numberOfComments--
        },
        handleLikeAdded(post) {
            post.numberOfLikes++
            post.isLiked = true
        },
        handleLikeDeleted(post) {
            post.numberOfLikes--
            post.isLiked = false
        }
    }
}
</script>

<template>
    <div>
        <ul class="list-group">
            <li class="list-group-item" v-for="photo in postsData.records" :key="photo.photoId">
                <PostItem @post-deleted="handlePostDeleted" @comment-added="handleCommentAdded(photo)"
                    @comment-deleted="handleCommentDeleted(photo)" @like-added="handleLikeAdded(photo)"
                    @like-deleted="handleLikeDeleted(photo)" :post-data="photo" />
            </li>
        </ul>
        <div class="text-center">
            <button v-if="postsData.hasNext" @click="$emit('load-more')" class="btn btn-primary mt-3 mb-3">Load
                More</button>
            <div v-if="!postsData.hasNext" class="alert alert-secondary mt-3 mb-3" role="alert">
                No more photos to show.
            </div>
        </div>
    </div>
</template>

<style scoped></style>