
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="图片预览:" prop="image">
    <el-input v-model="formData.image" :clearable="true" placeholder="请输入图片预览" />
</el-form-item>
        <el-form-item label="标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入标题" />
</el-form-item>
        <el-form-item label="跳转链接:" prop="url">
    <el-input v-model="formData.url" :clearable="true" placeholder="请输入跳转链接" />
</el-form-item>
        <el-form-item label="排序序号:" prop="orderId">
    <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入排序序号" />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCarouselImage,
  updateCarouselImage,
  findCarouselImage
} from '@/api/carouse/carouselImage'

defineOptions({
    name: 'CarouselImageForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const statusOptions = ref([])
const formData = ref({
            image: '',
            title: '',
            url: '',
            orderId: undefined,
            status: '',
        })
// 验证规则
const rule = reactive({
               image : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               url : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               orderId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCarouselImage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createCarouselImage(formData.value)
               break
             case 'update':
               res = await updateCarouselImage(formData.value)
               break
             default:
               res = await createCarouselImage(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
