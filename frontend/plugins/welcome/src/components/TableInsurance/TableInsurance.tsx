import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntInsurance } from '../../api/models/EntInsurance';
import moment from 'moment';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 
 const [insurances, setInsurances] = useState<EntInsurance[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
  const getInsurances = async () => {
    const res = await api.listInsurance({ limit: 10, offset: 0 });
    setLoading(false);
    setInsurances(res);
    console.log(res);
  };
  getInsurances();

}, [loading]);
 
 const deleteInsurances = async (id: number) => {
   const res = await api.deleteInsurance({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
            <TableCell align="center">ลำดับ Order</TableCell>
            <TableCell align="center">Car No.xx (หมายเลขประจำรถ)</TableCell>
            <TableCell align="center">ผู้รับผิดชอบประจำรถ</TableCell>
            <TableCell align="center">สถานะรถพยาบาล</TableCell>
            <TableCell align="center">เจ้าหน้าที่จัดการรถที่บันทึก</TableCell>
            <TableCell align="center">เวลาที่บันทึก</TableCell>
            <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {insurances === undefined
            ? null
            : insurances.map((item) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.car?.carNo}</TableCell>
             <TableCell align="center">{item.edges?.staff?.staffName}</TableCell>
             <TableCell align="center">{item.edges?.statuscar?.statusDetail}</TableCell>
             <TableCell align="center">{item.edges?.user?.userEmail}</TableCell>
             <TableCell align="center">{item.edges?.user?.userEmail}</TableCell>
             <TableCell align="center">{moment(item.addedTime).format('yyyy-MM-ddThh:mm')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   if (item.id === undefined){
                     return;
                   }
                   deleteInsurances(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
