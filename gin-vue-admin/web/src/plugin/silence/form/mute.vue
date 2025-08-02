
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="被禁言用户ID:" prop="userId">
          <el-input v-model="formData.userId" :clearable="true"  placeholder="请输入被禁言用户ID" />
       </el-form-item>
        <el-form-item label="禁言范围:" prop="muteType">
           <el-select v-model="formData.muteType" placeholder="请选择禁言范围" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in mutetypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="禁言开始时间:" prop="startTime">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="禁言结束时间:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="禁言原因:" prop="reason">
           <el-select v-model="formData.reason" placeholder="请选择禁言原因" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in reasonOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="操作人ID:" prop="operatorId">
          <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="禁言状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择禁言状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in statuOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="禁言天数:" prop="muteDay">
           <el-select v-model="formData.muteDay" placeholder="请选择禁言天数" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in speechOptions" :key="key" :label="item.label" :value="item.value" />
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
  createMute,
  updateMute,
  findMute
} from '@/plugin/silence/api/mute'

defineOptions({
    name: 'MuteForm'
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
const mutetypeOptions = ref([])
const reasonOptions = ref([])
const statuOptions = ref([])
const speechOptions = ref([])
const formData = ref({
            userId: '',
            muteType: '',
            startTime: new Date(),
            endTime: new Date(),
            reason: '',
            operatorId: undefined,
            status: '',
            muteDay: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               muteType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               endTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               reason : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               operatorId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               muteDay : [{
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
      const res = await findMute({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    mutetypeOptions.value = await getDictFunc('mutetype')
    reasonOptions.value = await getDictFunc('reason')
    statuOptions.value = await getDictFunc('statu')
    speechOptions.value = await getDictFunc('speech')
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
               res = await createMute(formData.value)
               break
             case 'update':
               res = await updateMute(formData.value)
               break
             default:
               res = await createMute(formData.value)
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
