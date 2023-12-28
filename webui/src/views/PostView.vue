<script>
export default {
  data: function () {
    return {
      errormsg: null,
      loading: false,
      caption: "",
      photoSelected: false,
      fileError: null
    }
  },
  methods: {
    async upload() {
      this.loading = true
      this.errormsg = null
      try {
        const formData = new FormData()
        formData.append("caption", this.caption)
        formData.append("photo", this.$refs.fileInput.files[0])
        await this.$axios.post("/photos/", formData)
        window.location.reload();
      } catch (e) {
        if (e.response && e.response.status === 400) {
          this.errormsg = "Invalid username."
        } else if (e.response && e.response.status === 401) {
          this.errormsg = "Unauthorized."
        } else if (e.response && e.response.status === 413) {
          this.errormsg = "Photo size too large, please upload a photo smaller than 10 MB."
        } else if (e.response && e.response.status === 415) {
          this.errormsg = "Wrong photo format, please upload a png photo."
        } else if (e.response && e.response.status === 500) {
          this.errormsg = "An internal error occurred. Please try again later."
        } else {
          this.errormsg = e.toString()
        }
      } finally {
        this.loading = false
      }
    },
    isCaptionValid() {
      return this.caption.length < 2200
    },
    validateFile() {
      const fileInput = this.$refs.fileInput
      this.fileError = (fileInput.files.length === 0 || !fileInput.files[0].name.toLowerCase().endsWith(".png"))
      this.photoSelected = !this.fileError
    },
  },
  mounted() {
    this.$emit("logged-in")
  },
}
</script>

<template>
  <div>
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h2>Post</h2>
    </div>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <form @submit.prevent="upload">
      <div class="mb-3">
        <label for="formFile" class="form-label">Photo</label>
        <input ref="fileInput" class="form-control" type="file" id="formFile" accept=".png" @change="validateFile"
          :class="{ 'is-invalid': fileError }">
        <div v-if="fileError" class="invalid-feedback">Wrong file format, please upload a photo in png format.</div>
        <label for="captionArea" class="form-label">Caption</label>
        <textarea class="form-control" id="captionArea" rows="3" v-model="caption"
          :class="{ 'is-invalid': !isCaptionValid() }"></textarea>
        <div v-if="!isCaptionValid()" class="invalid-feedback">Caption must be less than 2200 characters.</div>
      </div>
      <!-- Button trigger modal -->
      <button :disabled="!photoSelected || !isCaptionValid() || loading" type="button" class="btn btn-primary"
        data-bs-toggle="modal" data-bs-target="#uploadConfirmationModal">
        Post
      </button>
      <!-- Modal -->
      <div class="modal fade" id="uploadConfirmationModal" tabindex="-1" role="dialog"
        aria-labelledby="uploadConfirmationModalLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="uploadConfirmationModalLabel">Confirmation</h5>
              <button type="button" class="close" data-bs-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              Are you sure you want to upload the post?
              <LoadingSpinner :loading="loading"></LoadingSpinner>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
              <button type="submit" class="btn btn-primary" data-bs-dismiss="modal">Upload</button>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<style scoped></style>
