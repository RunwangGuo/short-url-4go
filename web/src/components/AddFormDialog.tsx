import * as React from 'react';
import {
	Button,
	Dialog,
	DialogActions,
	DialogContent,
	DialogContentText,
	DialogTitle,
	TextField,
} from '@mui/material';
import {split, trim} from 'lodash';
import SnackAlert from './SnackAlert';
import useService from '../service/service';
import {useAlert} from '../hooks';

type Props = {
	visible: boolean;
	onOk?: () => void;
	onCancel?: () => void;
};

export default function AddFormDialog(props: Props) {
	const {visible, onOk, onCancel} = props;

	const {alertVisible, alertMessage, alertColor, showAlert, closeAlert} = useAlert();
	const {generate, generating} = useService();

	const handleDialogClose = (_event: Record<string, unknown>, reason: 'backdropClick' | 'escapeKeyDown') => {
		if (reason === 'backdropClick') {
			return;
		}

		handleClose();
	};

	const handleClose = () => {
		if (onCancel) {
			onCancel();
		}
	};

	const handleSubmit = (data: Record<string, string>) => {
		const {urls, token} = data;
		if (!trim(urls).length) {
			showAlert('请填写正确的链接', 'error');
			return;
		}

		const strings = split(urls, '\n')
			.map((s) => trim(s))
			.filter((s) => s);
		if (!trim(token).length) {
			showAlert('请填写正确的安全码', 'error');
			return;
		}

		generate({
			token,
			urls: strings,
		})
			.then(() => {
				showAlert('添加成功', 'success');
				if (onOk) {
					onOk();
				}
			})
			.catch((err) => {
				showAlert(err.toString(), 'error');
			});
	};

	return (
		<div className={'dialog-container'}>
			<Dialog
				open={visible}
				disableEscapeKeyDown={true}
				PaperProps={{
					component: 'form',
					onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
						event.preventDefault();
						const formData = new FormData(event.currentTarget);
						const formJson = Object.fromEntries((formData as any).entries());
						handleSubmit(formJson);
					},
				}}
				onClose={handleDialogClose}
			>
				<DialogTitle>添加链接</DialogTitle>
				<DialogContent>
					<DialogContentText>输入链接后点击确定将生成对应短链接</DialogContentText>
					<TextField
						name={'urls'}
						type={'urls'}
						label={'链接'}
						margin={'dense'}
						variant={'standard'}
						autoFocus={true}
						required={true}
						fullWidth={true}
						multiline={true}
						rows={4}
						placeholder={'请输入链接，多个链接按行分隔'}
					/>
					<TextField
						name={'token'}
						type={'text'}
						label={'安全码'}
						margin={'dense'}
						variant={'standard'}
						required={true}
						placeholder={'请输入安全码'}
					/>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose}>取消</Button>
					<Button type={'submit'} disabled={generating}>
            确定
					</Button>
				</DialogActions>
			</Dialog>
			<SnackAlert
				visible={alertVisible}
				message={alertMessage}
				color={alertColor}
				onClose={closeAlert}
			/>
		</div>
	);
}
