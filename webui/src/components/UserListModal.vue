<script>
export default {
    emits: ["updateUsers", "resetUsers"],
    props: ["usersData", "title"],
    data() {
        return {
            modalId: crypto.randomUUID()
        }
    }
}
</script>

<template>
    <button @click="$emit('update-users')" type="button" class="btn btn-primary btn-sm me-2" data-bs-toggle="modal"
        :data-bs-target="'#usersModal' + modalId">
        <svg class="feather">
            <use href="/feather-sprite-v4.29.0.svg#eye" />
        </svg>
    </button>

    <div class="modal fade" :id="'usersModal' + modalId" tabindex="-1" role="dialog"
        :aria-labelledby="'usersModalLabel' + modalId" aria-hidden="true">
        <div class="modal-dialog" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" :id="'usersModalLabel' + modalId"> {{ title }} </h5>
                    <button @click="$emit('reset-users')" type="button" class="close" data-bs-dismiss="modal"
                        aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                </div>
                <div class="modal-body">
                    <div class="mt-4">
                        <ul class="list-group">
                            <a class="list-group-item list-group-item-action" v-for="user in usersData.records"
                                :key="user.userId">
                                <UserItem :user-data="user" />
                            </a>
                        </ul>
                        <button v-if="usersData.records !== null && usersData.hasNext" @click="$emit('update-users')"
                            class="btn btn-primary mt-3 mb-3">Load
                            More</button>
                        <div v-if="!usersData.hasNext" class="alert alert-secondary mt-3 mb-3" role="alert">
                            No more users to show.
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style></style>