
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="违规原因:" prop="violationMean">
    <el-select v-model="formData.violationMean" placeholder="请选择违规原因" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in violation_meanOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="处理结果:" prop="result">
    <el-input v-model="formData.result" :clearable="true" placeholder="请输入处理结果" />
</el-form-item>
        <el-form-item label="处理时间:" prop="resultTime">
    <el-date-picker v-model="formData.resultTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createViolation,
  updateViolation,
  findViolation
} from '@/api/violation/violation'

defineOptions({
    name: 'ViolationForm'
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
const violation_meanOptions = ref([])
const formData = ref({
            violationMean: '',
            result: '',
            resultTime: new Date(),
        })
// 验证规则
const rule = reactive({
               violationMean : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               result : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               resultTime : [{
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
      const res = await findViolation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    violation_meanOptions.value = await getDictFunc('violation_mean')
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
               res = await createViolation(formData.value)
               break
             case 'update':
               res = await updateViolation(formData.value)
               break
             default:
               res = await createViolation(formData.value)
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
