import React, { FC, useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core'
import { DefaultApi } from '../../api/apis';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { EntProduct } from '../../api/models/EntProduct';
import { EntMember } from '../../api/models/EntMember';
import { EntCategory } from '../../api/models/EntCategory';
import { EntOfficer } from '../../api/models/EntOfficer';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import InputAdornment from '@material-ui/core/InputAdornment';
import { Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';
import { ContentHeader } from '@backstage/core';
import { Link } from '@material-ui/core';
import { Theme, createStyles } from '@material-ui/core/styles';
import OutlinedInput from '@material-ui/core/OutlinedInput';


// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

/*interface inquiry {
  productid: string;
  memberid:   string;
  categoryid:   string;
  officerid:        string;
  inquiry_age:       number;
  inquiry_phone:      number;
  inquiry_name:      number;
  inquiry_messages: string;
  inquiryTimeMessages: string;
  // create_by: number;
}*/

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [alerttype, setAlertType] = useState(String);

  const [products, setProducts] = useState<EntProduct[]>([]);
  const [members, setMembers] = useState<EntMember[]>([]);
  const [categorys, setCategorys] = useState<EntCategory[]>([]);
  const [officers, setOfficers] = useState<EntOfficer[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);


  const [errorAgeMessages, setErrorAgemessages] = React.useState('');
  const [errorNameMessages, setErrorNamemessages] = React.useState('');
  const [errorPhoneMessages, setErrorPhonemessages] = React.useState('');
  const [errorMessages, setErrorMessages] = React.useState('');

  const [productid, setProductid] = useState(Number);
  const [memberid, setMemberid] = useState(Number);
  const [categoryid, setCategoryid] = useState(Number);
  const [officerid, setOfficerid] = useState(Number);
  const [inquiry_age_messages, setInquiryAgeMessages] = useState(Number);
  const [inquiry_messages, setInquiryMessages] = useState(String);
  const [inquiry_phone_messages, setInquiryPhoneMessages] = useState(String);
  const [inquiry_name_messages, setInquiryNameMessages] = useState(String);
  //const [inquiryTimeMessages, setInquiryTimeMessages] = useState(String);

  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const setErrorMessege = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: 'บันทึกข้อมูลไม่สำเร็จ',
      text: title,
    });
  }

  useEffect(() => {

    const getProducts = async () => {

      const p = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(false);
      setProducts(p);
    };
    getProducts();

    const getMembers = async () => {

      const u = await api.listMember({});
      setLoading(false);
      setMembers(u);
    };
    getMembers();

    const getCategorys = async () => {

      const cg = await api.listCategory({ limit: 10, offset: 0 });
      setLoading(false);
      setCategorys(cg);
    };
    getCategorys();

    const getOfficers = async () => {

      const of = await api.listOfficer({ limit: 10, offset: 0 });
      setLoading(false);
      setOfficers(of);
    };
    getOfficers();


    const data = localStorage.getItem("memberdata");
    if (data) {
      setMemberid(Number(localStorage.getItem("memberdata")));
      setLoading(false);
    }

    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("positiondata")));
      setLoading(false);
      if (jobdata != "สมาชิกระบบประกันสุขภาพ") {
        localStorage.setItem("memberdata", JSON.stringify(null));
        localStorage.setItem("positiondata", JSON.stringify(null));
        history.pushState("", "", "./SignIn");
        window.location.reload(false);
      }
      else {
        setMemberid(Number(localStorage.getItem("memberdata")))
      }
    }
    checkJobPosition();

  }, [loading]);

  const product_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProductid(event.target.value as number);
  };

  const member_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setMemberid(event.target.value as number);
  };

  const category_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setCategoryid(event.target.value as number);
  };

  const officer_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setOfficerid(event.target.value as number);
  };

  const inquiryAgeMessages_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('inquiry_age_messages', validateValue)
    setInquiryAgeMessages(event.target.value as number);
  };

  const inquiryPhoneMessages_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('inquiry_phone_messages', validateValue)
    setInquiryPhoneMessages(event.target.value as string);
  };

  const inquiryNameMessages_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('inquiry_name_messages', validateValue)
    setInquiryNameMessages(event.target.value as string);
  };

  const inquiryMessages_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('inquiry_messages', validateValue)
    setInquiryMessages(event.target.value as string);
  };

  const ValidateNamemessages = (val: string) => {
    return val.match("^[a-zA-Z ]+$");
  }

  const ValidateAgemessages = (val: number) => {
    return val >= 1 && val <= 999 ? true : false
  }

  const ValidatePhonemessages = (val: string) => {
    return val.match("^[0-9]{10}$");
  }

  const ValidateMessages = (val: string) => {
    return val.match("^[ก-๙0-9a-zA-Z- ./\\s]+$");
  }


  const checkpattern = (id: string, value: string) => {
    console.log(value);
    switch (id) {
      case 'inquiry_name_messages':
        ValidateNamemessages(value) ? setErrorNamemessages('') : setErrorNamemessages('ระบุชื่อเป็นภาษาอังกฤษให้ถูกต้อง และตัวแรกเป็นตัวพิมพ์ใหญ่');
        return;

      case 'inquiry_age_messages':
        ValidateAgemessages(Number(value)) ? setErrorAgemessages('') : setErrorAgemessages('กรุณากรอกอายุให้ถูกต้อง');
        return;
      case 'inquiry_phone_messages':
        ValidatePhonemessages(value) ? setErrorPhonemessages('') : setErrorPhonemessages('กรุณากรอกเบอร์โทรศัพท์ให้ถูกต้อง');
        return;
      case 'inquiry_messages':
        ValidateMessages(value) ? setErrorMessages('') : setErrorMessages('กรุณากรอกรายละเอียด');
        return;
      default:
        return;
    }
  }

  const CreateInquiry = async () => {
    /*if ((productid != null) && (categoryid != null) && (officerid != null) && (inquiry_messages != null) && (inquiry_messages != "") && (inquiry_phone_messages != null) && (inquiry_phone_messages != "")) */
     {
      const apiUrl = 'http://localhost:8080/api/v1/inquirys';
      const inquiry = {

        inquiryAgeMessages: Number(inquiry_age_messages),
        inquiryNameMessages: inquiry_name_messages,
        inquiryMessages: inquiry_messages,
        inquiryPhoneMessages: inquiry_phone_messages,
        inquiryTimeMessages: "",
        productID: productid,
        memberID: memberid,
        categoryID: categoryid,
        officerID: officerid,
      };
      console.log(inquiry);

      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(inquiry),
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            Toast.fire({
              icon: 'success',
              title: 'บันทึกข้อมูลสำเร็จ',
            });
            setLoading(true);
          } else {
            if (data.error.Name === undefined) {
              ErrorCaseCheck(data.error);
            } else {
              ErrorCaseCheck(data.error.Name);
            }
          }
        });
    };
  };
  const ErrorCaseCheck = (field: string) => {
    switch (field) {
      case 'Inquiry_phone_messages':
        setErrorMessege("error", "ระบุเลขเบอร์โทรศัพท์ให้ครบ 10 หลัก");
        return;
      case 'Inquiry_age_messages':
        setErrorMessege("error", "ระบุอายุให้ถูกต้อง และต้องมีอายุ 1 ปีขึ้นไป");
        return;
      case 'Inquiry_name_messages':
        setErrorMessege("error", "ระบุชื่อเป็นภาษาอังกฤษให้ถูกต้อง และตัวแรกเป็นตัวพิมพ์ใหญ่");
        return;
      case 'Inquiry_messages':
        setErrorMessege("error", "ระบุรายละเอียด");
        return;
      default:
        //setErrorMessege("error", "บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }
  /*const inquiryTimeMessages_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setInquiryTimeMessages(event.target.value as string);
  };*/

  return (
    <Page theme={pageTheme.app}>
      <Header style={HeaderCustom} title={`ระบบติดต่อสอบถามข้อมูล`}>
      </Header>
      <Content>
        <ContentHeader title="บันทึกการสอบถามข้อมูล">

        <Link component={RouterLink} to="/inquirysearch">
           <Button variant="contained" color="secondary">
             ประวัติการสอบถามข้อมูล
           </Button>
         </Link>

        </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>สมาชิกระบบประกันสุขภาพ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl
                className={classes.formControl}
                variant="outlined"
              >
                <TextField
                  id="user"
                  //label="ชื่อสมาชิก"
                  type="string"
                  size="medium"
                  variant="outlined"
                  value={members.filter((filter: EntMember) => filter.id == memberid).map((item: EntMember) => `${item.memberEmail}`)}
                  style={{ width: 300 }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อสมาชิก</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error={errorNameMessages ? true : false}
                  id="outlined-number"
                  name="inquiry_name_messages"
                  multiline
                  rows={1}
                  value={inquiry_name_messages}
                  type="string"
                  onChange={inquiryNameMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText={errorNameMessages}
                />
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>อายุ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error={errorAgeMessages ? true : false}
                  id="outlined-number"
                  name="inquiry_age_messages"
                  multiline
                  rows={1}
                  value={inquiry_age_messages}
                  type="string"
                  onChange={inquiryAgeMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText={errorAgeMessages}
                />
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เบอร์โทรติดต่อกลับ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error={errorPhoneMessages ? true : false}
                  id="outlined-number"
                  name="inquiry_phone_messages"
                  multiline
                  rows={1}
                  value={inquiry_phone_messages}
                  onChange={inquiryPhoneMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText={errorPhoneMessages}
                />
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>รายการสอบถาม</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรายการสอบถาม</InputLabel>
                <Select
                  name="category"
                  value={categoryid || ''} // (undefined || '') = ''
                  onChange={category_id_handleChange}
                >
                  {categorys.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.categoryName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อผลิตภัณฑ์</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผลิตภัณฑ์</InputLabel>
                <Select
                  name="product"
                  value={productid || ''} // (undefined || '') = ''
                  onChange={product_id_handleChange}
                >
                  {products.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.productName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>พนักงานบริษัทประกันสุขภาพ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงานบริษัทประกันสุขภาพ</InputLabel>
                <Select
                  name="officer"
                  value={officerid || ''} // (undefined || '') = ''
                  onChange={officer_id_handleChange}
                >
                  {officers.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.officerName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>รายละเอียด</div>
            </Grid>
            <Grid item xs={8}>
            <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  //style={{ width: 300 }}
                  error={errorMessages ? true : false}
                  id="outlined-number"
                  name="inquiry_messages"
                  multiline
                  rows={3}
                  value={inquiry_messages}
                  onChange={inquiryMessages_handleChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  variant="outlined"
                  helperText={errorMessages}
                />
              </FormControl>
            </Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreateInquiry();
                }}
              >
                บันทึกการสอบถาม
                </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};