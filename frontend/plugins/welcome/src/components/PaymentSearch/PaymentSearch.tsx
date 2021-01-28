import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
} from '@material-ui/core'
import SearchIcon from '@material-ui/icons/Search';
import TextField from '@material-ui/core/TextField';
import { Link, Link as RouterLink } from 'react-router-dom';
//import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
/* import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select'; */
import moment from 'moment';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
/* import { EntUser } from '../../api/models/EntUser';
import { EntDistance } from '../../api/models/EntDistance';
import { EntUrgent } from '../../api/models/EntUrgent'; */
//import { EntUser } from '../../api/models/EntUser'; 
import { EntPayment } from '../../api/models/EntPayment';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },

    withoutLabel: {
      marginTop: theme.spacing(1),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 100,
    },
    table: {
        minWidth: 650,
      },
  }),
);

/* const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'ระบบบันทึกการใช้รถพยาบาล' };
 */
const check = {
  customercheck : true
}

export default function Searchtable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [payments, setPayments] = useState<EntPayment[]>([]);
  const [loading, setLoading] = useState(true);
  const profile = { givenName: 'ระบบค้นหาการชำระเบี้ยประกัน' };
  const [status, setStatus] = useState(false);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [memberid, setMember] = useState(Number);
  
  const [paymentsearch, setPaymentSearch] = useState(String);



  useEffect(() => {

    const checkPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
      setLoading(false);
      if (jobdata != "สมาชิกระบบประกันสุขภาพ" ) {
        localStorage.setItem("memberdata",JSON.stringify(null));
        localStorage.setItem("positiondata",JSON.stringify(null));
        history.pushState("","","./");
        window.location.reload(false);        
      }
      else{
          setMember(Number(localStorage.getItem("memberdata")))
      }
    }
  checkPosition();
  }, [loading]);



  const SearchPayment = async () => {
    const res = await api.listPayment({ offset: 0 });
    const search = Paymentsearch(res);
    setErrorMessege("ไม่พบข้อมูลที่ค้นหา");
        setAlertType("error");
        setPayments([]);
        if(search.length > 0){
            Object.entries(check).map(([key, value]) =>{
                if (value == true){
                    setErrorMessege("พบข้อมูลที่ค้นหา");
                    setAlertType("success");
                    setPayments(search);
                }
            })
        }

        setStatus(true);
  }

  const Paymentsearch = (res: any) => {
    const data = res.filter((filter: EntPayment) => filter?.accountName?.includes(paymentsearch))
    if (data.length != 0 && paymentsearch != "") {
        return data;
    }
    else{
      return data;
        }
    }

  const handleSearchChange = (event: any) => {
    setPaymentSearch(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.library}>
     <Header
       title={`${profile.givenName || ':)'}`}
     >
   </Header>
     <Content>
        <ContentHeader title="ค้นหาการชำระเบี้ยประกัน">
        {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
        </ContentHeader>
        <Container maxWidth="lg">
        <Grid container spacing={0}>
          <Grid item xs={12}></Grid>
          <Grid item xs={1}>
              <div className={classes.paper}>ชื่อ-สุกล</div>
          </Grid>
          <Grid item xs={3}>
          <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  style={{ width: 300 }}
                  id="payment"
                  label=""
                  variant="outlined"
                  color="secondary"
                  type="string"
                  size="small"
                  value={paymentsearch}
                  onChange={handleSearchChange}
                  name="insurance_identification"
                  multiline
                  rows={1}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
          </Grid>
          <Grid item xs={4}>
          <Button
            onClick={() => {
              SearchPayment();
            }}
            variant="contained"
            color="secondary"
            style={{ width: 90 }}
            startIcon={<SearchIcon />}
          >
            ค้นหา
          </Button>
         </Grid>
        </Grid>
      </Container>
      <br></br>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">ลำดับ</TableCell>
                <TableCell align="center">ชื่อ-สกุล</TableCell>
                <TableCell align="center">หมายเลขกรมธรรม์</TableCell>
                <TableCell align="center">จำนวนเบี้ยประกัน</TableCell>
                <TableCell align="center">วัน-เวลาที่ชำระเงิน</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {payments.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell> 
                  <TableCell align="center">{item.accountName}</TableCell> 
                  <TableCell align="center">{item.edges?.insurance?.id}</TableCell>
                  <TableCell align="center">{item.price}</TableCell>
                  <TableCell align="center">{moment(item.transferTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>
      </Content>
    </Page>
  );
};