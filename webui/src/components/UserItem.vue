<script>
export default {
    props: ["userData"]
}
</script>

<template>
    <RouterLink :to="`/${userData.username}`" class="nav-link">
        <svg class="feather">
            <use href="/feather-sprite-v4.29.0.svg#user" />
        </svg>
        {{ userData.username }}
    </RouterLink>
</template>

<style></style>