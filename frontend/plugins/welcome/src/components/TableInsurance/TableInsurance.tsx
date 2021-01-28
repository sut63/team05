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
import { EntMember } from '../../api/models/EntMember';
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
 const [members, setMembers] = useState<EntMember[]>([]);
 const [memberid, setMemberID] = useState(Number);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
  const getInsurances = async () => {
    const res = await api.listInsurance({ limit: 10, offset: 0 });
    setLoading(false);
    setInsurances(res);
    console.log(res);
  };
  getInsurances();

  const checkJobPosition = async () => {
    const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
    setLoading(false);
    if (jobdata != "สมาชิกระบบประกันสุขภาพ" ) {
      localStorage.setItem("memberdata",JSON.stringify(null));
      localStorage.setItem("positiondata",JSON.stringify(null));
      history.pushState("","","./SignIn");
      window.location.reload(false);        
    }
    else{
        setMemberID(Number(localStorage.getItem("memberdata")))
    }
  }
checkJobPosition();

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
            <TableCell align="center">ลำดับประกันที่ซื้อ</TableCell>
            <TableCell align="center">ชื่อ Prodcut</TableCell>
            <TableCell align="center">ชื่อผู้รับผลประโยชน์</TableCell>
            <TableCell align="center">ราคา</TableCell>
            <TableCell align="center">โรงพยาบาล</TableCell>
            <TableCell align="center">วันเวลาที่ซื้อ</TableCell>
            <TableCell align="center">วันที่ต้องการจ่ายงวดแรก</TableCell>
            <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {insurances === undefined
            ? null
            : insurances.map((item) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.product?.productName}</TableCell>
             <TableCell align="center">{item.insuranceInsurer}</TableCell>
             <TableCell align="center">{item.edges?.product?.productPrice}</TableCell>
             <TableCell align="center">{item.edges?.hospital?.hospitalName}</TableCell>
             <TableCell align="center">{moment(item.insuranceTimeBuy).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">{moment(item.insuranceTimeFirstpay).format('DD/MM/YYYY')}</TableCell>
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
