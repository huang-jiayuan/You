
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="w-[380px]"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="silence_Mute" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="silence_Mute" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="silence_Mute" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="被禁言用户ID" prop="userId" width="120" />

            <el-table-column align="left" label="禁言范围" prop="muteType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.muteType,mutetypeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="禁言开始时间" prop="startTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="禁言结束时间" prop="endTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="禁言原因" prop="reason" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.reason,reasonOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="操作人ID" prop="operatorId" width="120" />

            <el-table-column align="left" label="禁言状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,statuOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="禁言天数" prop="muteDay" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.muteDay,speechOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateMuteFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="被禁言用户ID:" prop="userId">
    <el-input v-model="formData.userId" :clearable="true" placeholder="请输入被禁言用户ID" />
</el-form-item>
             <el-form-item label="禁言范围:" prop="muteType">
    <el-select v-model="formData.muteType" placeholder="请选择禁言范围" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in mutetypeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="禁言开始时间:" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="禁言结束时间:" prop="endTime">
    <el-date-picker v-model="formData.endTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
             <el-form-item label="禁言原因:" prop="reason">
    <el-select v-model="formData.reason" placeholder="请选择禁言原因" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in reasonOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="操作人ID:" prop="operatorId">
    <el-input v-model.number="formData.operatorId" :clearable="true" placeholder="请输入操作人ID" />
</el-form-item>
             <el-form-item label="禁言状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择禁言状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in statuOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="禁言天数:" prop="muteDay">
    <el-select v-model="formData.muteDay" placeholder="请选择禁言天数" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in speechOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="被禁言用户ID">
    {{ detailFrom.userId }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言范围">
    {{ detailFrom.muteType }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言开始时间">
    {{ detailFrom.startTime }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言结束时间">
    {{ detailFrom.endTime }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言原因">
    {{ detailFrom.reason }}
</el-descriptions-item>
                 <el-descriptions-item label="操作人ID">
    {{ detailFrom.operatorId }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言状态">
    {{ detailFrom.status }}
</el-descriptions-item>
                 <el-descriptions-item label="禁言天数">
    {{ detailFrom.muteDay }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createMute,
  deleteMute,
  deleteMuteByIds,
  updateMute,
  findMute,
  getMuteList
} from '@/plugin/silence/api/mute'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Mute'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               muteType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               startTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               endTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               reason : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               operatorId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               muteDay : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getMuteList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    mutetypeOptions.value = await getDictFunc('mutetype')
    reasonOptions.value = await getDictFunc('reason')
    statuOptions.value = await getDictFunc('statu')
    speechOptions.value = await getDictFunc('speech')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteMuteFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteMuteByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateMuteFunc = async(row) => {
    const res = await findMute({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteMuteFunc = async (row) => {
    const res = await deleteMute({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        userId: '',
        muteType: '',
        startTime: new Date(),
        endTime: new Date(),
        reason: '',
        operatorId: undefined,
        status: '',
        muteDay: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findMute({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
