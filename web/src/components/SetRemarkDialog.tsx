import * as React from 'react';
import SnackAlert from './SnackAlert';
import {useAlert} from '../hooks';
import {Button, Dialog, DialogActions, DialogContent, DialogTitle, TextField} from '@mui/material';
import {trim} from 'lodash';
import useService from '../service/service';

type Props = {
	visible: boolean;
	targets: string[];
	defaultValue?: string;
	onOk?: () => void;
	onCancel?: () => void;
};

export default function setRemarkDialog(props: Props) {
	const {visible, targets, defaultValue, onOk, onCancel} = props;
	const {alertVisible, alertMessage, alertColor, showAlert, closeAlert} = useAlert();
	const {remark, remarking} = useService();

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
		const {token, content} = data;
		if (!trim(token).length) {
			showAlert('请填写正确的安全码', 'error');
			return;
		}

		remark({
			token,
			targets,
			remark: content,
		})
			.then(() => {
				showAlert('操作成功', 'success');
				if (onOk) {
					onOk();
				}

				handleClose();
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
				<DialogTitle>备注</DialogTitle>
				<DialogContent>
					<TextField
						name={'content'}
						type={'text'}
						label={'备注'}
						margin={'dense'}
						variant={'standard'}
						autoFocus={true}
						fullWidth={true}
						multiline={true}
						rows={4}
						placeholder={'请输入备注'}
						defaultValue={defaultValue ?? ''}
					/>
					<TextField
						name={'token'}
						type={'text'}
						label={'安全码'}
						margin={'dense'}
						variant={'standard'}
						autoFocus={true}
						required={true}
						placeholder={'请输入安全码'}
					/>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose}>取消</Button>
					<Button type={'submit'} disabled={remarking}>
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
