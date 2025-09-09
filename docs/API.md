# Kết quả xổ số API


API phục vụ việc lấy danh sách tỉnh, tra cứu kết quả xổ số và kiểm tra vé số trúng thưởng. Được viết bằng **Golang + Gin**.

---

## Chạy API Server
```bash
./KetQuaXoSo --api
```

---

## Sử dụng
- Link API Gốc: `http://localhost:8080/api`

---

## Endpoints
- Lấy danh sách tỉnh `GET /api/province`:
bash:
```bash
curl -X GET http://localhost:8080/api/province
```
output: 
```json
["Miền Bắc","An Giang","Bình Dương","Bình Định","Bạc Liêu","Bình Phước","Bến Tre","Bình Thuận","Cà Mau","Cần Thơ","Đắk Lắk","Đồng Nai","Đà Nẵng","Đắk Nông","Đồng Tháp","Gia Lai","Hồ Chí Minh","Hậu Giang","Kiên Giang","Khánh Hòa","Kon Tum","Long An","Lâm Đồng","Ninh Thuận","Phú Yên","Quảng Bình","Quảng Ngãi","Quảng Nam","Quảng Trị","Sóc Trăng","Tiền Giang","Tây Ninh","Thừa Thiên Huế","Trà Vinh","Vĩnh Long","Vũng Tàu"]
```
- Tra cứu kết quả xổ số `POST /api/check`:
bash:
```bash
curl -X POST http://localhost:8080/api/check -H "Content-Type: application/json" -d '{"province":"Quảng Trị"}'
```
output:
```json
{"province":"Quảng Trị","results":[{"Province":"Quảng Trị","Date":"04/09","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 04/09 (Thứ Năm)","Prizes":{"1":["27518"],"2":["11510"],"3":["31548","80246"],"4":["36329","65314","21353","55983","33863","80469","07183"],"5":["5712"],"6":["9919","5202","3134"],"7":["571"],"8":["59"],"ĐB":["477536"]}},{"Province":"Quảng Trị","Date":"28/08","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 28/08 (Thứ Năm)","Prizes":{"1":["72176"],"2":["51500"],"3":["75352","68104"],"4":["14125","84713","79107","64130","54584","80787","27054"],"5":["3406"],"6":["2968","8884","9136"],"7":["002"],"8":["21"],"ĐB":["690290"]}},{"Province":"Quảng Trị","Date":"21/08","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 21/08 (Thứ Năm)","Prizes":{"1":["47737"],"2":["72978"],"3":["31923","90276"],"4":["64499","81353","12187","95969","86989","10500","03546"],"5":["7900"],"6":["4851","0935","4209"],"7":["692"],"8":["89"],"ĐB":["240200"]}},{"Province":"Quảng Trị","Date":"14/08","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 14/08 (Thứ Năm)","Prizes":{"1":["94473"],"2":["49849"],"3":["42778","38079"],"4":["42106","56886","33775","27670","11349","86000","75008"],"5":["1948"],"6":["6359","6268","6135"],"7":["317"],"8":["75"],"ĐB":["939537"]}},{"Province":"Quảng Trị","Date":"07/08","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 07/08 (Thứ Năm)","Prizes":{"1":["20657"],"2":["87799"],"3":["87697","99411"],"4":["50254","88317","58966","85334","53371","99366","56608"],"5":["6857"],"6":["3826","6482","8503"],"7":["883"],"8":["36"],"ĐB":["950528"]}},{"Province":"Quảng Trị","Date":"31/07","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 31/07 (Thứ Năm)","Prizes":{"1":["65725"],"2":["50319"],"3":["28095","45269"],"4":["91264","38210","47458","89125","27467","65458","13964"],"5":["8872"],"6":["8182","4694","4762"],"7":["813"],"8":["98"],"ĐB":["618532"]}},{"Province":"Quảng Trị","Date":"24/07","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 24/07 (Thứ Năm)","Prizes":{"1":["59451"],"2":["71132"],"3":["46540","50622"],"4":["57094","34315","99853","43499","06670","45192","13052"],"5":["1591"],"6":["0841","1913","1610"],"7":["237"],"8":["04"],"ĐB":["018860"]}},{"Province":"Quảng Trị","Date":"17/07","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 17/07 (Thứ Năm)","Prizes":{"1":["02391"],"2":["41139"],"3":["94450","29365"],"4":["77262","35601","33873","42318","86294","46225","61906"],"5":["4305"],"6":["6392","5260","4447"],"7":["642"],"8":["79"],"ĐB":["647894"]}},{"Province":"Quảng Trị","Date":"10/07","Title":"KẾT QUẢ XỔ SỐ QUẢNG TRỊ NGÀY 10/07 (Thứ Năm)","Prizes":{"1":["41352"],"2":["67848"],"3":["24696","00802"],"4":["89271","48469","51593","42554","12889","93559","40902"],"5":["3255"],"6":["1645","9574","7105"],"7":["620"],"8":["18"],"ĐB":["975817"]}}]}
```
- Kiểm tra vé số `POST /api/check-ticket`:
bash:
```bash
curl -X POST http://localhost:8080/api/check-ticket -H "Content-Type: application/json" -d '{"province":"Quảng Trị","date":"04/09","number":"477536"}'
```
output:
```json
{"province":"Quảng Trị","date":"04/09","input":"477536","prize":"ĐB","match":"477536"}
```
