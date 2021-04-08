<template>
    <div class="spu-sku">
        <div class="wrap">
            <h2>配置区域</h2>
            <div class="spu-sku-item" v-for="(item,index) in spus" :key="index">
                <div class="spu">
                    <strong>spu:</strong>
                    <el-input placeholder="SPU" v-model="item.name"></el-input>
                    <el-button class="close" plain circle @click="handleDelSpu(index)">
                        <i class="el-icon-close"></i>
                    </el-button>
                </div>
                <div class="sku">
                    <strong>sku:</strong>
                    <span class="sku-item" v-for="(it,i) in item.skus" :key="i">
                        <el-input placeholder="SKU" v-model="it.name" @change="handleSkuChange"></el-input>
                        <el-button class="close" plain circle @click="handleDelSku(item, index, i)">
                            <i class="el-icon-close"></i>
                        </el-button>
                    </span>
                    <el-button @click="handleAddSku(item)">添加规格</el-button>
                </div>
            </div>
            <h2>
                <el-button @click="handleAddSpu">添加规格项</el-button>
            </h2>
        </div>
        <div class="list">
            <el-table :data="list" :span-method="spanMethod">
                <template v-for="(item, index) in spus">
                    <el-table-column :key="index" :label="item.name" :prop="item.name" width="120px"></el-table-column>
                </template>
                <el-table-column label="其他参数" ></el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script>
let currentIndex = 0
const GenerateList = (spus, nestedIndex, result, row) => {
  result = result || []
  let spu, skus, spusLen
  spu = spus[nestedIndex]
  skus = spu.skus
  spusLen = spus.length
  row = row || {}
  row.skuNameArr = row.skuNameArr || []
  if (skus && skus.length) {
    for (let i = 0, len = skus.length; i < len; i++) {
      if (currentIndex === spusLen - 1) {
        let endIndex = nestedIndex === row.skuNameArr.length ? row.skuNameArr.length : nestedIndex - row.skuNameArr.length
        row.skuNameArr = row.skuNameArr.slice(0, endIndex)
      }

      let sku = skus[i]
      row[spu.name] = sku.name
      row.skuNameArr.push(`${spu.name}:${sku.name}`)
      if (nestedIndex < spusLen - 1) {
        GenerateList(spus, nestedIndex + 1, result, row)
      } else {
        row.name = row.skuNameArr.join(',')
        result.push(Object.assign({}, row))
        row.name = ''
        currentIndex = nestedIndex
      }
    }
  }
}
const generateSpanArr = (spus) => {
  const result = {}
  for (let i = 0, len = spus.length; i < len; i++) {
    const key = spus[i].name
    let pos = 0
    for (let j = 0, len = spus.length; j < len; j++) {
      result[key] = []
      if (j === 0) {
        result[key].push(1)
      } else {
        // 判断当前元素与上一个元素是否相同
        if (spus[j][key] === spus[j - 1][key]) {
          result[key][pos] += 1
          result[key].push(0)
        } else {
          result[key].push(1)
          pos = i
        }
      }
    }
  }
  return result
}
export default {
  data () {
    return {
      spus: [
        {
          name: '颜色',
          skus: [
            {
              name: '红色'
            },
            {
              name: '黄色'
            },
            {
              name: '绿色'
            }
          ]
        },
        {
          name: '尺寸',
          skus: [
            {
              name: '大'
            },
            {
              name: '小'
            }
          ]
        },
        {
          name: '容量',
          skus: [
            {
              name: '10ml'
            },
            {
              name: '20ml'
            }
          ]
        }
      ],
      list: []
    }
  },
  methods: {
    handleAddSpu () {
      this.spus.push({
        name: '',
        skus: [
          {
            name: ''
          }
        ]
      })
    },
    handleDelSpu (index) {
      this.spus.splice(index, 1)
    },
    handleAddSku (spu) {
      spu.skus.push({ name: '' })
    },
    handleDelSku (spu, spuIndex, index) {
      if (index === 0) {
        this.spus.splice(index, 1)
        return
      }
      spu.skus.splice(index, 1)
    },
    handleSkuChange (value) {
      // GenerateList(this.)
    },
    spanMethod ({ row, column, rowIndex, columnIndex }) {
      // if (columnIndex < this.spus.length) {
      //   // console.log(row)
      //   return {
      //     rowspan: 2,
      //     colspan: 1
      //   }
      // }
    }
  },
  mounted () {
    let result = []
    GenerateList(this.spus, 0, result)
    this.list = result
    console.log(generateSpanArr(this.spus))
  }
}
</script>

<style lang="scss" scoped>
    $-padding: 10px;
    $-margin: 10px;
    $-bgc: #f5f5f5;
    .el-input {
        width: 200px;
    }

    .spu-sku {
        .spu-sku-item {
            & + .spu-sku-item {
                margin-top: $-margin;
            }
        }

        .wrap {
            background-color: $-bgc;
            padding: $-padding;

            strong {
                margin-right: $-margin;
            }

            .close {
                padding: 6px;
                margin-left: $-margin;
            }

            .spu {
                margin-bottom: $-margin;
            }

            .sku {
                &-item {
                    margin-right: $-margin;
                }
            }
        }
    }
</style>
